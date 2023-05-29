//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnConfig) DeepCopyInto(out *AddOnConfig) {
	*out = *in
	out.ConfigGroupResource = in.ConfigGroupResource
	out.ConfigReferent = in.ConfigReferent
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnConfig.
func (in *AddOnConfig) DeepCopy() *AddOnConfig {
	if in == nil {
		return nil
	}
	out := new(AddOnConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnDeploymentConfig) DeepCopyInto(out *AddOnDeploymentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnDeploymentConfig.
func (in *AddOnDeploymentConfig) DeepCopy() *AddOnDeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(AddOnDeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnDeploymentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnDeploymentConfigList) DeepCopyInto(out *AddOnDeploymentConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddOnDeploymentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnDeploymentConfigList.
func (in *AddOnDeploymentConfigList) DeepCopy() *AddOnDeploymentConfigList {
	if in == nil {
		return nil
	}
	out := new(AddOnDeploymentConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnDeploymentConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnDeploymentConfigSpec) DeepCopyInto(out *AddOnDeploymentConfigSpec) {
	*out = *in
	if in.CustomizedVariables != nil {
		in, out := &in.CustomizedVariables, &out.CustomizedVariables
		*out = make([]CustomizedVariable, len(*in))
		copy(*out, *in)
	}
	if in.NodePlacement != nil {
		in, out := &in.NodePlacement, &out.NodePlacement
		*out = new(NodePlacement)
		(*in).DeepCopyInto(*out)
	}
	if in.Registries != nil {
		in, out := &in.Registries, &out.Registries
		*out = make([]ImageMirror, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnDeploymentConfigSpec.
func (in *AddOnDeploymentConfigSpec) DeepCopy() *AddOnDeploymentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AddOnDeploymentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnMeta) DeepCopyInto(out *AddOnMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnMeta.
func (in *AddOnMeta) DeepCopy() *AddOnMeta {
	if in == nil {
		return nil
	}
	out := new(AddOnMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnTemplate) DeepCopyInto(out *AddOnTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnTemplate.
func (in *AddOnTemplate) DeepCopy() *AddOnTemplate {
	if in == nil {
		return nil
	}
	out := new(AddOnTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnTemplateList) DeepCopyInto(out *AddOnTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddOnTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnTemplateList.
func (in *AddOnTemplateList) DeepCopy() *AddOnTemplateList {
	if in == nil {
		return nil
	}
	out := new(AddOnTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnTemplateSpec) DeepCopyInto(out *AddOnTemplateSpec) {
	*out = *in
	in.AgentSpec.DeepCopyInto(&out.AgentSpec)
	if in.Registration != nil {
		in, out := &in.Registration, &out.Registration
		*out = make([]RegistrationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnTemplateSpec.
func (in *AddOnTemplateSpec) DeepCopy() *AddOnTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(AddOnTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagementAddOn) DeepCopyInto(out *ClusterManagementAddOn) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagementAddOn.
func (in *ClusterManagementAddOn) DeepCopy() *ClusterManagementAddOn {
	if in == nil {
		return nil
	}
	out := new(ClusterManagementAddOn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagementAddOn) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagementAddOnList) DeepCopyInto(out *ClusterManagementAddOnList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterManagementAddOn, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagementAddOnList.
func (in *ClusterManagementAddOnList) DeepCopy() *ClusterManagementAddOnList {
	if in == nil {
		return nil
	}
	out := new(ClusterManagementAddOnList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterManagementAddOnList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagementAddOnSpec) DeepCopyInto(out *ClusterManagementAddOnSpec) {
	*out = *in
	out.AddOnMeta = in.AddOnMeta
	out.AddOnConfiguration = in.AddOnConfiguration
	if in.SupportedConfigs != nil {
		in, out := &in.SupportedConfigs, &out.SupportedConfigs
		*out = make([]ConfigMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.InstallStrategy.DeepCopyInto(&out.InstallStrategy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagementAddOnSpec.
func (in *ClusterManagementAddOnSpec) DeepCopy() *ClusterManagementAddOnSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterManagementAddOnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagementAddOnStatus) DeepCopyInto(out *ClusterManagementAddOnStatus) {
	*out = *in
	if in.DefaultConfigReferences != nil {
		in, out := &in.DefaultConfigReferences, &out.DefaultConfigReferences
		*out = make([]DefaultConfigReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstallProgressions != nil {
		in, out := &in.InstallProgressions, &out.InstallProgressions
		*out = make([]InstallProgression, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagementAddOnStatus.
func (in *ClusterManagementAddOnStatus) DeepCopy() *ClusterManagementAddOnStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterManagementAddOnStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigCoordinates) DeepCopyInto(out *ConfigCoordinates) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigCoordinates.
func (in *ConfigCoordinates) DeepCopy() *ConfigCoordinates {
	if in == nil {
		return nil
	}
	out := new(ConfigCoordinates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigGroupResource) DeepCopyInto(out *ConfigGroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigGroupResource.
func (in *ConfigGroupResource) DeepCopy() *ConfigGroupResource {
	if in == nil {
		return nil
	}
	out := new(ConfigGroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMeta) DeepCopyInto(out *ConfigMeta) {
	*out = *in
	out.ConfigGroupResource = in.ConfigGroupResource
	if in.DefaultConfig != nil {
		in, out := &in.DefaultConfig, &out.DefaultConfig
		*out = new(ConfigReferent)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMeta.
func (in *ConfigMeta) DeepCopy() *ConfigMeta {
	if in == nil {
		return nil
	}
	out := new(ConfigMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigReference) DeepCopyInto(out *ConfigReference) {
	*out = *in
	out.ConfigGroupResource = in.ConfigGroupResource
	out.ConfigReferent = in.ConfigReferent
	if in.DesiredConfig != nil {
		in, out := &in.DesiredConfig, &out.DesiredConfig
		*out = new(ConfigSpecHash)
		**out = **in
	}
	if in.LastAppliedConfig != nil {
		in, out := &in.LastAppliedConfig, &out.LastAppliedConfig
		*out = new(ConfigSpecHash)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigReference.
func (in *ConfigReference) DeepCopy() *ConfigReference {
	if in == nil {
		return nil
	}
	out := new(ConfigReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigReferent) DeepCopyInto(out *ConfigReferent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigReferent.
func (in *ConfigReferent) DeepCopy() *ConfigReferent {
	if in == nil {
		return nil
	}
	out := new(ConfigReferent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpecHash) DeepCopyInto(out *ConfigSpecHash) {
	*out = *in
	out.ConfigReferent = in.ConfigReferent
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpecHash.
func (in *ConfigSpecHash) DeepCopy() *ConfigSpecHash {
	if in == nil {
		return nil
	}
	out := new(ConfigSpecHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSignerRegistrationConfig) DeepCopyInto(out *CustomSignerRegistrationConfig) {
	*out = *in
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(Subject)
		(*in).DeepCopyInto(*out)
	}
	out.SigningCA = in.SigningCA
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSignerRegistrationConfig.
func (in *CustomSignerRegistrationConfig) DeepCopy() *CustomSignerRegistrationConfig {
	if in == nil {
		return nil
	}
	out := new(CustomSignerRegistrationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedVariable) DeepCopyInto(out *CustomizedVariable) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedVariable.
func (in *CustomizedVariable) DeepCopy() *CustomizedVariable {
	if in == nil {
		return nil
	}
	out := new(CustomizedVariable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultConfigReference) DeepCopyInto(out *DefaultConfigReference) {
	*out = *in
	out.ConfigGroupResource = in.ConfigGroupResource
	if in.DesiredConfig != nil {
		in, out := &in.DesiredConfig, &out.DesiredConfig
		*out = new(ConfigSpecHash)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultConfigReference.
func (in *DefaultConfigReference) DeepCopy() *DefaultConfigReference {
	if in == nil {
		return nil
	}
	out := new(DefaultConfigReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubPermissionConfig) DeepCopyInto(out *HubPermissionConfig) {
	*out = *in
	out.RoleRef = in.RoleRef
	if in.SingleNamespace != nil {
		in, out := &in.SingleNamespace, &out.SingleNamespace
		*out = new(SingleNamespaceBindingConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubPermissionConfig.
func (in *HubPermissionConfig) DeepCopy() *HubPermissionConfig {
	if in == nil {
		return nil
	}
	out := new(HubPermissionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageMirror) DeepCopyInto(out *ImageMirror) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageMirror.
func (in *ImageMirror) DeepCopy() *ImageMirror {
	if in == nil {
		return nil
	}
	out := new(ImageMirror)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallConfigReference) DeepCopyInto(out *InstallConfigReference) {
	*out = *in
	out.ConfigGroupResource = in.ConfigGroupResource
	if in.DesiredConfig != nil {
		in, out := &in.DesiredConfig, &out.DesiredConfig
		*out = new(ConfigSpecHash)
		**out = **in
	}
	if in.LastKnownGoodConfig != nil {
		in, out := &in.LastKnownGoodConfig, &out.LastKnownGoodConfig
		*out = new(ConfigSpecHash)
		**out = **in
	}
	if in.LastAppliedConfig != nil {
		in, out := &in.LastAppliedConfig, &out.LastAppliedConfig
		*out = new(ConfigSpecHash)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallConfigReference.
func (in *InstallConfigReference) DeepCopy() *InstallConfigReference {
	if in == nil {
		return nil
	}
	out := new(InstallConfigReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallProgression) DeepCopyInto(out *InstallProgression) {
	*out = *in
	out.PlacementRef = in.PlacementRef
	if in.ConfigReferences != nil {
		in, out := &in.ConfigReferences, &out.ConfigReferences
		*out = make([]InstallConfigReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallProgression.
func (in *InstallProgression) DeepCopy() *InstallProgression {
	if in == nil {
		return nil
	}
	out := new(InstallProgression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallStrategy) DeepCopyInto(out *InstallStrategy) {
	*out = *in
	if in.Placements != nil {
		in, out := &in.Placements, &out.Placements
		*out = make([]PlacementStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallStrategy.
func (in *InstallStrategy) DeepCopy() *InstallStrategy {
	if in == nil {
		return nil
	}
	out := new(InstallStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeClientRegistrationConfig) DeepCopyInto(out *KubeClientRegistrationConfig) {
	*out = *in
	if in.HubPermissions != nil {
		in, out := &in.HubPermissions, &out.HubPermissions
		*out = make([]HubPermissionConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeClientRegistrationConfig.
func (in *KubeClientRegistrationConfig) DeepCopy() *KubeClientRegistrationConfig {
	if in == nil {
		return nil
	}
	out := new(KubeClientRegistrationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterAddOn) DeepCopyInto(out *ManagedClusterAddOn) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterAddOn.
func (in *ManagedClusterAddOn) DeepCopy() *ManagedClusterAddOn {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterAddOn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterAddOn) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterAddOnList) DeepCopyInto(out *ManagedClusterAddOnList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedClusterAddOn, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterAddOnList.
func (in *ManagedClusterAddOnList) DeepCopy() *ManagedClusterAddOnList {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterAddOnList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterAddOnList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterAddOnSpec) DeepCopyInto(out *ManagedClusterAddOnSpec) {
	*out = *in
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make([]AddOnConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterAddOnSpec.
func (in *ManagedClusterAddOnSpec) DeepCopy() *ManagedClusterAddOnSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterAddOnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterAddOnStatus) DeepCopyInto(out *ManagedClusterAddOnStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]ObjectReference, len(*in))
		copy(*out, *in)
	}
	out.AddOnMeta = in.AddOnMeta
	out.AddOnConfiguration = in.AddOnConfiguration
	if in.SupportedConfigs != nil {
		in, out := &in.SupportedConfigs, &out.SupportedConfigs
		*out = make([]ConfigGroupResource, len(*in))
		copy(*out, *in)
	}
	if in.ConfigReferences != nil {
		in, out := &in.ConfigReferences, &out.ConfigReferences
		*out = make([]ConfigReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Registrations != nil {
		in, out := &in.Registrations, &out.Registrations
		*out = make([]RegistrationConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.HealthCheck = in.HealthCheck
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterAddOnStatus.
func (in *ManagedClusterAddOnStatus) DeepCopy() *ManagedClusterAddOnStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterAddOnStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePlacement) DeepCopyInto(out *NodePlacement) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePlacement.
func (in *NodePlacement) DeepCopy() *NodePlacement {
	if in == nil {
		return nil
	}
	out := new(NodePlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRef) DeepCopyInto(out *PlacementRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRef.
func (in *PlacementRef) DeepCopy() *PlacementRef {
	if in == nil {
		return nil
	}
	out := new(PlacementRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementStrategy) DeepCopyInto(out *PlacementStrategy) {
	*out = *in
	out.PlacementRef = in.PlacementRef
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make([]AddOnConfig, len(*in))
		copy(*out, *in)
	}
	in.RolloutStrategy.DeepCopyInto(&out.RolloutStrategy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementStrategy.
func (in *PlacementStrategy) DeepCopy() *PlacementStrategy {
	if in == nil {
		return nil
	}
	out := new(PlacementStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrationConfig) DeepCopyInto(out *RegistrationConfig) {
	*out = *in
	in.Subject.DeepCopyInto(&out.Subject)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrationConfig.
func (in *RegistrationConfig) DeepCopy() *RegistrationConfig {
	if in == nil {
		return nil
	}
	out := new(RegistrationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrationSpec) DeepCopyInto(out *RegistrationSpec) {
	*out = *in
	if in.KubeClient != nil {
		in, out := &in.KubeClient, &out.KubeClient
		*out = new(KubeClientRegistrationConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomSigner != nil {
		in, out := &in.CustomSigner, &out.CustomSigner
		*out = new(CustomSignerRegistrationConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrationSpec.
func (in *RegistrationSpec) DeepCopy() *RegistrationSpec {
	if in == nil {
		return nil
	}
	out := new(RegistrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdate) DeepCopyInto(out *RollingUpdate) {
	*out = *in
	out.MaxConcurrency = in.MaxConcurrency
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdate.
func (in *RollingUpdate) DeepCopy() *RollingUpdate {
	if in == nil {
		return nil
	}
	out := new(RollingUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateWithCanary) DeepCopyInto(out *RollingUpdateWithCanary) {
	*out = *in
	out.Placement = in.Placement
	out.RollingUpdate = in.RollingUpdate
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateWithCanary.
func (in *RollingUpdateWithCanary) DeepCopy() *RollingUpdateWithCanary {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateWithCanary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStrategy) DeepCopyInto(out *RolloutStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdate)
		**out = **in
	}
	if in.RollingUpdateWithCanary != nil {
		in, out := &in.RollingUpdateWithCanary, &out.RollingUpdateWithCanary
		*out = new(RollingUpdateWithCanary)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStrategy.
func (in *RolloutStrategy) DeepCopy() *RolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(RolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningCARef) DeepCopyInto(out *SigningCARef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningCARef.
func (in *SigningCARef) DeepCopy() *SigningCARef {
	if in == nil {
		return nil
	}
	out := new(SigningCARef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleNamespaceBindingConfig) DeepCopyInto(out *SingleNamespaceBindingConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleNamespaceBindingConfig.
func (in *SingleNamespaceBindingConfig) DeepCopy() *SingleNamespaceBindingConfig {
	if in == nil {
		return nil
	}
	out := new(SingleNamespaceBindingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OrganizationUnits != nil {
		in, out := &in.OrganizationUnits, &out.OrganizationUnits
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
	in.DeepCopyInto(out)
	return out
}
