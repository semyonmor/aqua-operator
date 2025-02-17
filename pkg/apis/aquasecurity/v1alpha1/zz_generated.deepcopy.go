//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by operator-sdk-old. DO NOT EDIT.

package v1alpha1

import (
	operatorv1alpha1 "github.com/aquasecurity/aqua-operator/pkg/apis/operator/v1alpha1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AquaStarboard) DeepCopyInto(out *AquaStarboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AquaStarboard.
func (in *AquaStarboard) DeepCopy() *AquaStarboard {
	if in == nil {
		return nil
	}
	out := new(AquaStarboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AquaStarboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AquaStarboardList) DeepCopyInto(out *AquaStarboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AquaStarboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AquaStarboardList.
func (in *AquaStarboardList) DeepCopy() *AquaStarboardList {
	if in == nil {
		return nil
	}
	out := new(AquaStarboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AquaStarboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AquaStarboardSpec) DeepCopyInto(out *AquaStarboardSpec) {
	*out = *in
	if in.Infrastructure != nil {
		in, out := &in.Infrastructure, &out.Infrastructure
		*out = new(operatorv1alpha1.AquaInfrastructure)
		**out = **in
	}
	if in.StarboardService != nil {
		in, out := &in.StarboardService, &out.StarboardService
		*out = new(operatorv1alpha1.AquaService)
		(*in).DeepCopyInto(*out)
	}
	out.Config = in.Config
	if in.RegistryData != nil {
		in, out := &in.RegistryData, &out.RegistryData
		*out = new(operatorv1alpha1.AquaDockerRegistry)
		**out = **in
	}
	if in.ImageData != nil {
		in, out := &in.ImageData, &out.ImageData
		*out = new(operatorv1alpha1.AquaImage)
		**out = **in
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AquaStarboardSpec.
func (in *AquaStarboardSpec) DeepCopy() *AquaStarboardSpec {
	if in == nil {
		return nil
	}
	out := new(AquaStarboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AquaStarboardStatus) DeepCopyInto(out *AquaStarboardStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AquaStarboardStatus.
func (in *AquaStarboardStatus) DeepCopy() *AquaStarboardStatus {
	if in == nil {
		return nil
	}
	out := new(AquaStarboardStatus)
	in.DeepCopyInto(out)
	return out
}
