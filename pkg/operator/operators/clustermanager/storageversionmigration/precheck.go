package storageversionmigration

import (
	"context"
	"strings"
	"time"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	migrationv1alpha1 "sigs.k8s.io/kube-storage-version-migrator/pkg/apis/migration/v1alpha1"
	migrationv1alpha1client "sigs.k8s.io/kube-storage-version-migrator/pkg/clients/clientset/typed/migration/v1alpha1"
)

var checkList []string = []string{
	"managedclustersets.v1beta1.cluster.open-cluster-management.io",
	"managedclustersetbindings.v1beta1.cluster.open-cluster-management.io",
}

// Purpose: This function is used to solve a MCE 2.4 upgrade to 2.5 cases. In some MCE 2.4 environments,
// the managedclusterset and managedclustersetbinding CRDs still have both v1beta1 and v1beta2 versions
// in etcd and the privous migrations are failed of processing the CRDs.
//
// If we upgrade directly from MCE 2.4 to 2.5, the operator will apply the New CRDs(v1beta1 removed)
// and that will cause crash because the old version CRs still exist in etcd.
//
// So before the operator starts, we need to use "patch" to clean the conditions of the old migrations to trigger
// another migration job.
func MigrationPreCheck(ctx context.Context, config *rest.Config) error {
	klog.Info("Starting Migration Pre Check.")
	var err error

	// Check if storageversionmigration CRD exists.
	// If not, return nil.
	apiExtensionClient, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return err
	}
	_, err = apiExtensionClient.ApiextensionsV1().CustomResourceDefinitions().Get(ctx, "storageversionmigrations.migration.k8s.io", metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}
	klog.Info("StorageVersionMigration CRD exists.")

	migrationClient, err := migrationv1alpha1client.NewForConfig(config)
	if err != nil {
		return err
	}

	return wait.PollUntilContextTimeout(ctx, 3*time.Second, time.Minute, true, func(ctx context.Context) (bool, error) {
		// Get all storageversionmigration CRs.
		migrations, err := migrationClient.StorageVersionMigrations().List(ctx, metav1.ListOptions{})
		if err != nil {
			return false, err
		}

		// Filter the migrations which need to be updated.
		migrationsNeedUpdate := filterMigrations(migrations.Items, checkList)

		// If no migrations need to be update, return nil.
		if len(migrationsNeedUpdate) == 0 {
			return true, nil
		}

		// Patch migrations that need to be updated.
		for _, migration := range migrationsNeedUpdate {
			_, err := migrationClient.StorageVersionMigrations().Patch(ctx, migration.Name, types.JSONPatchType,
				[]byte(`[{"op":"remove", "path":"/status/conditions"}]`), metav1.PatchOptions{}, "status")
			if err != nil {
				return false, err
			}
			klog.Infof("Patched the migration %s.", migration.Name)
		}
		return false, nil
	})
}

func filterMigrations(migrations []migrationv1alpha1.StorageVersionMigration, checkList []string) []migrationv1alpha1.StorageVersionMigration {
	var migrationsNeedUpdate []migrationv1alpha1.StorageVersionMigration
	for _, migration := range migrations {
		// only check the migrations in the checkList.
		if strings.Contains(strings.Join(checkList, ","), migration.Name) {
			// If migration succeed, skip.
			if checkMigrationStatus(migration) {
				continue
			}
			migrationsNeedUpdate = append(migrationsNeedUpdate, migration)
		}
	}
	return migrationsNeedUpdate
}

func checkMigrationStatus(migration migrationv1alpha1.StorageVersionMigration) bool {
	// The "Succeeded" condition is expected exist and be true.
	for _, condition := range migration.Status.Conditions {
		if condition.Type == "Succeeded" && condition.Status == "True" {
			return true
		}
	}
	return false
}
