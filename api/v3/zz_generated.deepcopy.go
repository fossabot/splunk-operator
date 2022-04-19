//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v3

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentContext) DeepCopyInto(out *AppDeploymentContext) {
	*out = *in
	in.AppFrameworkConfig.DeepCopyInto(&out.AppFrameworkConfig)
	if in.AppsSrcDeployStatus != nil {
		in, out := &in.AppsSrcDeployStatus, &out.AppsSrcDeployStatus
		*out = make(map[string]AppSrcDeployInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentContext.
func (in *AppDeploymentContext) DeepCopy() *AppDeploymentContext {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentInfo) DeepCopyInto(out *AppDeploymentInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentInfo.
func (in *AppDeploymentInfo) DeepCopy() *AppDeploymentInfo {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppFrameworkSpec) DeepCopyInto(out *AppFrameworkSpec) {
	*out = *in
	out.Defaults = in.Defaults
	if in.VolList != nil {
		in, out := &in.VolList, &out.VolList
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.AppSources != nil {
		in, out := &in.AppSources, &out.AppSources
		*out = make([]AppSourceSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppFrameworkSpec.
func (in *AppFrameworkSpec) DeepCopy() *AppFrameworkSpec {
	if in == nil {
		return nil
	}
	out := new(AppFrameworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSourceDefaultSpec) DeepCopyInto(out *AppSourceDefaultSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSourceDefaultSpec.
func (in *AppSourceDefaultSpec) DeepCopy() *AppSourceDefaultSpec {
	if in == nil {
		return nil
	}
	out := new(AppSourceDefaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSourceSpec) DeepCopyInto(out *AppSourceSpec) {
	*out = *in
	out.AppSourceDefaultSpec = in.AppSourceDefaultSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSourceSpec.
func (in *AppSourceSpec) DeepCopy() *AppSourceSpec {
	if in == nil {
		return nil
	}
	out := new(AppSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSrcDeployInfo) DeepCopyInto(out *AppSrcDeployInfo) {
	*out = *in
	if in.AppDeploymentInfoList != nil {
		in, out := &in.AppDeploymentInfoList, &out.AppDeploymentInfoList
		*out = make([]AppDeploymentInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSrcDeployInfo.
func (in *AppSrcDeployInfo) DeepCopy() *AppSrcDeployInfo {
	if in == nil {
		return nil
	}
	out := new(AppSrcDeployInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundlePushInfo) DeepCopyInto(out *BundlePushInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundlePushInfo.
func (in *BundlePushInfo) DeepCopy() *BundlePushInfo {
	if in == nil {
		return nil
	}
	out := new(BundlePushInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheManagerSpec) DeepCopyInto(out *CacheManagerSpec) {
	*out = *in
	out.IndexAndCacheManagerCommonSpec = in.IndexAndCacheManagerCommonSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheManagerSpec.
func (in *CacheManagerSpec) DeepCopy() *CacheManagerSpec {
	if in == nil {
		return nil
	}
	out := new(CacheManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMaster) DeepCopyInto(out *ClusterMaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMaster.
func (in *ClusterMaster) DeepCopy() *ClusterMaster {
	if in == nil {
		return nil
	}
	out := new(ClusterMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterMaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMasterList) DeepCopyInto(out *ClusterMasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterMaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMasterList.
func (in *ClusterMasterList) DeepCopy() *ClusterMasterList {
	if in == nil {
		return nil
	}
	out := new(ClusterMasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterMasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMasterSpec) DeepCopyInto(out *ClusterMasterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	in.AppFrameworkConfig.DeepCopyInto(&out.AppFrameworkConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMasterSpec.
func (in *ClusterMasterSpec) DeepCopy() *ClusterMasterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMasterStatus) DeepCopyInto(out *ClusterMasterStatus) {
	*out = *in
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	out.BundlePushTracker = in.BundlePushTracker
	if in.ResourceRevMap != nil {
		in, out := &in.ResourceRevMap, &out.ResourceRevMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AppContext.DeepCopyInto(&out.AppContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMasterStatus.
func (in *ClusterMasterStatus) DeepCopy() *ClusterMasterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterMasterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSplunkSpec) DeepCopyInto(out *CommonSplunkSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.EtcVolumeStorageConfig = in.EtcVolumeStorageConfig
	out.VarVolumeStorageConfig = in.VarVolumeStorageConfig
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ObsoleteLicenseManagerRef = in.ObsoleteLicenseManagerRef
	out.LicenseManagerRef = in.LicenseManagerRef
	out.ClusterMasterRef = in.ClusterMasterRef
	out.MonitoringConsoleRef = in.MonitoringConsoleRef
	if in.ExtraEnv != nil {
		in, out := &in.ExtraEnv, &out.ExtraEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSplunkSpec.
func (in *CommonSplunkSpec) DeepCopy() *CommonSplunkSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSplunkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexAndCacheManagerCommonSpec) DeepCopyInto(out *IndexAndCacheManagerCommonSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexAndCacheManagerCommonSpec.
func (in *IndexAndCacheManagerCommonSpec) DeepCopy() *IndexAndCacheManagerCommonSpec {
	if in == nil {
		return nil
	}
	out := new(IndexAndCacheManagerCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexAndGlobalCommonSpec) DeepCopyInto(out *IndexAndGlobalCommonSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexAndGlobalCommonSpec.
func (in *IndexAndGlobalCommonSpec) DeepCopy() *IndexAndGlobalCommonSpec {
	if in == nil {
		return nil
	}
	out := new(IndexAndGlobalCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexConfDefaultsSpec) DeepCopyInto(out *IndexConfDefaultsSpec) {
	*out = *in
	out.IndexAndGlobalCommonSpec = in.IndexAndGlobalCommonSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexConfDefaultsSpec.
func (in *IndexConfDefaultsSpec) DeepCopy() *IndexConfDefaultsSpec {
	if in == nil {
		return nil
	}
	out := new(IndexConfDefaultsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexSpec) DeepCopyInto(out *IndexSpec) {
	*out = *in
	out.IndexAndCacheManagerCommonSpec = in.IndexAndCacheManagerCommonSpec
	out.IndexAndGlobalCommonSpec = in.IndexAndGlobalCommonSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexSpec.
func (in *IndexSpec) DeepCopy() *IndexSpec {
	if in == nil {
		return nil
	}
	out := new(IndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerCluster) DeepCopyInto(out *IndexerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerCluster.
func (in *IndexerCluster) DeepCopy() *IndexerCluster {
	if in == nil {
		return nil
	}
	out := new(IndexerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterList) DeepCopyInto(out *IndexerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IndexerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterList.
func (in *IndexerClusterList) DeepCopy() *IndexerClusterList {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterMemberStatus) DeepCopyInto(out *IndexerClusterMemberStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterMemberStatus.
func (in *IndexerClusterMemberStatus) DeepCopy() *IndexerClusterMemberStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterSpec) DeepCopyInto(out *IndexerClusterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterSpec.
func (in *IndexerClusterSpec) DeepCopy() *IndexerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterStatus) DeepCopyInto(out *IndexerClusterStatus) {
	*out = *in
	if in.IndexerSecretChanged != nil {
		in, out := &in.IndexerSecretChanged, &out.IndexerSecretChanged
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.IdxcPasswordChangedSecrets != nil {
		in, out := &in.IdxcPasswordChangedSecrets, &out.IdxcPasswordChangedSecrets
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]IndexerClusterMemberStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterStatus.
func (in *IndexerClusterStatus) DeepCopy() *IndexerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManager) DeepCopyInto(out *LicenseManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManager.
func (in *LicenseManager) DeepCopy() *LicenseManager {
	if in == nil {
		return nil
	}
	out := new(LicenseManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManagerList) DeepCopyInto(out *LicenseManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManagerList.
func (in *LicenseManagerList) DeepCopy() *LicenseManagerList {
	if in == nil {
		return nil
	}
	out := new(LicenseManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManagerSpec) DeepCopyInto(out *LicenseManagerSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	in.AppFrameworkConfig.DeepCopyInto(&out.AppFrameworkConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManagerSpec.
func (in *LicenseManagerSpec) DeepCopy() *LicenseManagerSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseManagerStatus) DeepCopyInto(out *LicenseManagerStatus) {
	*out = *in
	in.AppContext.DeepCopyInto(&out.AppContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseManagerStatus.
func (in *LicenseManagerStatus) DeepCopy() *LicenseManagerStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMaster) DeepCopyInto(out *LicenseMaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMaster.
func (in *LicenseMaster) DeepCopy() *LicenseMaster {
	if in == nil {
		return nil
	}
	out := new(LicenseMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterList) DeepCopyInto(out *LicenseMasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseMaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterList.
func (in *LicenseMasterList) DeepCopy() *LicenseMasterList {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterSpec) DeepCopyInto(out *LicenseMasterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	in.AppFrameworkConfig.DeepCopyInto(&out.AppFrameworkConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterSpec.
func (in *LicenseMasterSpec) DeepCopy() *LicenseMasterSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterStatus) DeepCopyInto(out *LicenseMasterStatus) {
	*out = *in
	in.AppContext.DeepCopyInto(&out.AppContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterStatus.
func (in *LicenseMasterStatus) DeepCopy() *LicenseMasterStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsole) DeepCopyInto(out *MonitoringConsole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsole.
func (in *MonitoringConsole) DeepCopy() *MonitoringConsole {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringConsole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsoleList) DeepCopyInto(out *MonitoringConsoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonitoringConsole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsoleList.
func (in *MonitoringConsoleList) DeepCopy() *MonitoringConsoleList {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitoringConsoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsoleSpec) DeepCopyInto(out *MonitoringConsoleSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	in.AppFrameworkConfig.DeepCopyInto(&out.AppFrameworkConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsoleSpec.
func (in *MonitoringConsoleSpec) DeepCopy() *MonitoringConsoleSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConsoleStatus) DeepCopyInto(out *MonitoringConsoleStatus) {
	*out = *in
	out.BundlePushTracker = in.BundlePushTracker
	if in.ResourceRevMap != nil {
		in, out := &in.ResourceRevMap, &out.ResourceRevMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AppContext.DeepCopyInto(&out.AppContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConsoleStatus.
func (in *MonitoringConsoleStatus) DeepCopy() *MonitoringConsoleStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringConsoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadCluster) DeepCopyInto(out *SearchHeadCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadCluster.
func (in *SearchHeadCluster) DeepCopy() *SearchHeadCluster {
	if in == nil {
		return nil
	}
	out := new(SearchHeadCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterList) DeepCopyInto(out *SearchHeadClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SearchHeadCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterList.
func (in *SearchHeadClusterList) DeepCopy() *SearchHeadClusterList {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterMemberStatus) DeepCopyInto(out *SearchHeadClusterMemberStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterMemberStatus.
func (in *SearchHeadClusterMemberStatus) DeepCopy() *SearchHeadClusterMemberStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterSpec) DeepCopyInto(out *SearchHeadClusterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	in.AppFrameworkConfig.DeepCopyInto(&out.AppFrameworkConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterSpec.
func (in *SearchHeadClusterSpec) DeepCopy() *SearchHeadClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterStatus) DeepCopyInto(out *SearchHeadClusterStatus) {
	*out = *in
	if in.ShcSecretChanged != nil {
		in, out := &in.ShcSecretChanged, &out.ShcSecretChanged
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.AdminSecretChanged != nil {
		in, out := &in.AdminSecretChanged, &out.AdminSecretChanged
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.AdminPasswordChangedSecrets != nil {
		in, out := &in.AdminPasswordChangedSecrets, &out.AdminPasswordChangedSecrets
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]SearchHeadClusterMemberStatus, len(*in))
		copy(*out, *in)
	}
	in.AppContext.DeepCopyInto(&out.AppContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterStatus.
func (in *SearchHeadClusterStatus) DeepCopy() *SearchHeadClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartStoreSpec) DeepCopyInto(out *SmartStoreSpec) {
	*out = *in
	if in.VolList != nil {
		in, out := &in.VolList, &out.VolList
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.IndexList != nil {
		in, out := &in.IndexList, &out.IndexList
		*out = make([]IndexSpec, len(*in))
		copy(*out, *in)
	}
	out.Defaults = in.Defaults
	out.CacheManagerConf = in.CacheManagerConf
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartStoreSpec.
func (in *SmartStoreSpec) DeepCopy() *SmartStoreSpec {
	if in == nil {
		return nil
	}
	out := new(SmartStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spec) DeepCopyInto(out *Spec) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.ServiceTemplate.DeepCopyInto(&out.ServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spec.
func (in *Spec) DeepCopy() *Spec {
	if in == nil {
		return nil
	}
	out := new(Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Standalone) DeepCopyInto(out *Standalone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Standalone.
func (in *Standalone) DeepCopy() *Standalone {
	if in == nil {
		return nil
	}
	out := new(Standalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Standalone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneList) DeepCopyInto(out *StandaloneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Standalone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneList.
func (in *StandaloneList) DeepCopy() *StandaloneList {
	if in == nil {
		return nil
	}
	out := new(StandaloneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandaloneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneSpec) DeepCopyInto(out *StandaloneSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	in.AppFrameworkConfig.DeepCopyInto(&out.AppFrameworkConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneSpec.
func (in *StandaloneSpec) DeepCopy() *StandaloneSpec {
	if in == nil {
		return nil
	}
	out := new(StandaloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneStatus) DeepCopyInto(out *StandaloneStatus) {
	*out = *in
	in.SmartStore.DeepCopyInto(&out.SmartStore)
	if in.ResourceRevMap != nil {
		in, out := &in.ResourceRevMap, &out.ResourceRevMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AppContext.DeepCopyInto(&out.AppContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneStatus.
func (in *StandaloneStatus) DeepCopy() *StandaloneStatus {
	if in == nil {
		return nil
	}
	out := new(StandaloneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassSpec) DeepCopyInto(out *StorageClassSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassSpec.
func (in *StorageClassSpec) DeepCopy() *StorageClassSpec {
	if in == nil {
		return nil
	}
	out := new(StorageClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAndTypeSpec) DeepCopyInto(out *VolumeAndTypeSpec) {
	*out = *in
	out.VolumeSpec = in.VolumeSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAndTypeSpec.
func (in *VolumeAndTypeSpec) DeepCopy() *VolumeAndTypeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeAndTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
