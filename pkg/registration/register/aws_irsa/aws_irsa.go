package aws_irsa

import (
	"context"
	"fmt"

	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	clusterv1 "open-cluster-management.io/api/cluster/v1"
	operatorv1 "open-cluster-management.io/api/operator/v1"

	"open-cluster-management.io/ocm/pkg/common/helpers"
	"open-cluster-management.io/ocm/pkg/registration/register"
)

//TODO: Remove these constants in once we have the function fully implemented for the AWSIRSADriver

const (
	// TLSKeyFile is the name of tls key file in kubeconfigSecret
	TLSKeyFile = "tls.key"
	// TLSCertFile is the name of the tls cert file in kubeconfigSecret
	TLSCertFile                 = "tls.crt"
	ManagedClusterArn           = "managed-cluster-arn"
	ManagedClusterIAMRoleSuffix = "managed-cluster-iam-role-suffix"
)

type AWSIRSADriver struct {
	name                     string
	managedClusterArn        string
	hubClusterArn            string
	managedClusterRoleSuffix string
}

func (c *AWSIRSADriver) Process(
	ctx context.Context, controllerName string, secret *corev1.Secret, additionalSecretData map[string][]byte,
	recorder events.Recorder, opt any) (*corev1.Secret, *metav1.Condition, error) {

	awsOption, ok := opt.(*AWSOption)
	if !ok {
		return nil, nil, fmt.Errorf("option type is not correct")
	}

	isApproved, err := awsOption.AWSIRSAControl.isApproved(c.name)
	if err != nil {
		return nil, nil, err
	}
	if !isApproved {
		return nil, nil, nil
	}

	recorder.Eventf("EKSRegistrationRequestApproved", "An EKS registration request is approved for %s", controllerName)
	return secret, nil, nil
}

func (c *AWSIRSADriver) BuildKubeConfigFromTemplate(kubeConfig *clientcmdapi.Config) *clientcmdapi.Config {
	hubClusterAccountId, hubClusterName := helpers.GetAwsAccountIdAndClusterName(c.hubClusterArn)
	awsRegion := helpers.GetAwsRegion(c.hubClusterArn)
	kubeConfig.AuthInfos = map[string]*clientcmdapi.AuthInfo{register.DefaultKubeConfigAuth: {
		Exec: &clientcmdapi.ExecConfig{
			APIVersion: "client.authentication.k8s.io/v1beta1",
			Command:    "aws",
			Args: []string{
				"--region",
				awsRegion,
				"eks",
				"get-token",
				"--cluster-name",
				hubClusterName,
				"--output",
				"json",
				"--role",
				fmt.Sprintf("arn:aws:iam::%s:role/ocm-hub-%s", hubClusterAccountId, c.managedClusterRoleSuffix),
			},
		},
	}}
	return kubeConfig
}

func (c *AWSIRSADriver) InformerHandler(option any) (cache.SharedIndexInformer, factory.EventFilterFunc) {
	awsOption, ok := option.(*AWSOption)
	if !ok {
		utilruntime.Must(fmt.Errorf("option type is not correct"))
	}
	return awsOption.AWSIRSAControl.Informer(), awsOption.EventFilterFunc
}

func (c *AWSIRSADriver) IsHubKubeConfigValid(ctx context.Context, secretOption register.SecretOption) (bool, error) {
	// TODO: implement the logic to validate the kubeconfig
	return true, nil
}

func (c *AWSIRSADriver) ManagedClusterDecorator(cluster *clusterv1.ManagedCluster) *clusterv1.ManagedCluster {
	if cluster.Annotations == nil {
		cluster.Annotations = make(map[string]string)
	}
	cluster.Annotations[operatorv1.ClusterAnnotationsKeyPrefix+"/"+ManagedClusterArn] = c.managedClusterArn
	cluster.Annotations[operatorv1.ClusterAnnotationsKeyPrefix+"/"+ManagedClusterIAMRoleSuffix] = c.managedClusterRoleSuffix
	return cluster
}

func NewAWSIRSADriver(managedClusterArn string, managedClusterRoleSuffix string, hubClusterArn string, name string) register.RegisterDriver {
	return &AWSIRSADriver{
		managedClusterArn:        managedClusterArn,
		managedClusterRoleSuffix: managedClusterRoleSuffix,
		hubClusterArn:            hubClusterArn,
		name:                     name,
	}
}
