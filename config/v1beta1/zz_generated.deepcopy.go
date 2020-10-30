// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kube-scheduler/config/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extender) DeepCopyInto(out *Extender) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(v1.ExtenderTLSConfig)
		(*in).DeepCopyInto(*out)
	}
	out.HTTPTimeout = in.HTTPTimeout
	if in.ManagedResources != nil {
		in, out := &in.ManagedResources, &out.ManagedResources
		*out = make([]v1.ExtenderManagedResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extender.
func (in *Extender) DeepCopy() *Extender {
	if in == nil {
		return nil
	}
	out := new(Extender)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterPodAffinityArgs) DeepCopyInto(out *InterPodAffinityArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.HardPodAffinityWeight != nil {
		in, out := &in.HardPodAffinityWeight, &out.HardPodAffinityWeight
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterPodAffinityArgs.
func (in *InterPodAffinityArgs) DeepCopy() *InterPodAffinityArgs {
	if in == nil {
		return nil
	}
	out := new(InterPodAffinityArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterPodAffinityArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(int32)
		**out = **in
	}
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	out.ClientConnection = in.ClientConnection
	if in.HealthzBindAddress != nil {
		in, out := &in.HealthzBindAddress, &out.HealthzBindAddress
		*out = new(string)
		**out = **in
	}
	if in.MetricsBindAddress != nil {
		in, out := &in.MetricsBindAddress, &out.MetricsBindAddress
		*out = new(string)
		**out = **in
	}
	in.DebuggingConfiguration.DeepCopyInto(&out.DebuggingConfiguration)
	if in.PercentageOfNodesToScore != nil {
		in, out := &in.PercentageOfNodesToScore, &out.PercentageOfNodesToScore
		*out = new(int32)
		**out = **in
	}
	if in.PodInitialBackoffSeconds != nil {
		in, out := &in.PodInitialBackoffSeconds, &out.PodInitialBackoffSeconds
		*out = new(int64)
		**out = **in
	}
	if in.PodMaxBackoffSeconds != nil {
		in, out := &in.PodMaxBackoffSeconds, &out.PodMaxBackoffSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]KubeSchedulerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extenders != nil {
		in, out := &in.Extenders, &out.Extenders
		*out = make([]Extender, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerConfiguration.
func (in *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerProfile) DeepCopyInto(out *KubeSchedulerProfile) {
	*out = *in
	if in.SchedulerName != nil {
		in, out := &in.SchedulerName, &out.SchedulerName
		*out = new(string)
		**out = **in
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = new(Plugins)
		(*in).DeepCopyInto(*out)
	}
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerProfile.
func (in *KubeSchedulerProfile) DeepCopy() *KubeSchedulerProfile {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLabelArgs) DeepCopyInto(out *NodeLabelArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.PresentLabels != nil {
		in, out := &in.PresentLabels, &out.PresentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AbsentLabels != nil {
		in, out := &in.AbsentLabels, &out.AbsentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PresentLabelsPreference != nil {
		in, out := &in.PresentLabelsPreference, &out.PresentLabelsPreference
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AbsentLabelsPreference != nil {
		in, out := &in.AbsentLabelsPreference, &out.AbsentLabelsPreference
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLabelArgs.
func (in *NodeLabelArgs) DeepCopy() *NodeLabelArgs {
	if in == nil {
		return nil
	}
	out := new(NodeLabelArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeLabelArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesFitArgs) DeepCopyInto(out *NodeResourcesFitArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.IgnoredResources != nil {
		in, out := &in.IgnoredResources, &out.IgnoredResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnoredResourceGroups != nil {
		in, out := &in.IgnoredResourceGroups, &out.IgnoredResourceGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesFitArgs.
func (in *NodeResourcesFitArgs) DeepCopy() *NodeResourcesFitArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesFitArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesFitArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesLeastAllocatedArgs) DeepCopyInto(out *NodeResourcesLeastAllocatedArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesLeastAllocatedArgs.
func (in *NodeResourcesLeastAllocatedArgs) DeepCopy() *NodeResourcesLeastAllocatedArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesLeastAllocatedArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesLeastAllocatedArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesMostAllocatedArgs) DeepCopyInto(out *NodeResourcesMostAllocatedArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesMostAllocatedArgs.
func (in *NodeResourcesMostAllocatedArgs) DeepCopy() *NodeResourcesMostAllocatedArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesMostAllocatedArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesMostAllocatedArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	in.Args.DeepCopyInto(&out.Args)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginSet) DeepCopyInto(out *PluginSet) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginSet.
func (in *PluginSet) DeepCopy() *PluginSet {
	if in == nil {
		return nil
	}
	out := new(PluginSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugins) DeepCopyInto(out *Plugins) {
	*out = *in
	if in.QueueSort != nil {
		in, out := &in.QueueSort, &out.QueueSort
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreFilter != nil {
		in, out := &in.PreFilter, &out.PreFilter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PostFilter != nil {
		in, out := &in.PostFilter, &out.PostFilter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreScore != nil {
		in, out := &in.PreScore, &out.PreScore
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Reserve != nil {
		in, out := &in.Reserve, &out.Reserve
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Permit != nil {
		in, out := &in.Permit, &out.Permit
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreBind != nil {
		in, out := &in.PreBind, &out.PreBind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Bind != nil {
		in, out := &in.Bind, &out.Bind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PostBind != nil {
		in, out := &in.PostBind, &out.PostBind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugins.
func (in *Plugins) DeepCopy() *Plugins {
	if in == nil {
		return nil
	}
	out := new(Plugins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTopologySpreadArgs) DeepCopyInto(out *PodTopologySpreadArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DefaultConstraints != nil {
		in, out := &in.DefaultConstraints, &out.DefaultConstraints
		*out = make([]corev1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTopologySpreadArgs.
func (in *PodTopologySpreadArgs) DeepCopy() *PodTopologySpreadArgs {
	if in == nil {
		return nil
	}
	out := new(PodTopologySpreadArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodTopologySpreadArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedToCapacityRatioArgs) DeepCopyInto(out *RequestedToCapacityRatioArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = make([]UtilizationShapePoint, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedToCapacityRatioArgs.
func (in *RequestedToCapacityRatioArgs) DeepCopy() *RequestedToCapacityRatioArgs {
	if in == nil {
		return nil
	}
	out := new(RequestedToCapacityRatioArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RequestedToCapacityRatioArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAffinityArgs) DeepCopyInto(out *ServiceAffinityArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.AffinityLabels != nil {
		in, out := &in.AffinityLabels, &out.AffinityLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AntiAffinityLabelsPreference != nil {
		in, out := &in.AntiAffinityLabelsPreference, &out.AntiAffinityLabelsPreference
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAffinityArgs.
func (in *ServiceAffinityArgs) DeepCopy() *ServiceAffinityArgs {
	if in == nil {
		return nil
	}
	out := new(ServiceAffinityArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAffinityArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UtilizationShapePoint) DeepCopyInto(out *UtilizationShapePoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UtilizationShapePoint.
func (in *UtilizationShapePoint) DeepCopy() *UtilizationShapePoint {
	if in == nil {
		return nil
	}
	out := new(UtilizationShapePoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBindingArgs) DeepCopyInto(out *VolumeBindingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.BindTimeoutSeconds != nil {
		in, out := &in.BindTimeoutSeconds, &out.BindTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBindingArgs.
func (in *VolumeBindingArgs) DeepCopy() *VolumeBindingArgs {
	if in == nil {
		return nil
	}
	out := new(VolumeBindingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBindingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
