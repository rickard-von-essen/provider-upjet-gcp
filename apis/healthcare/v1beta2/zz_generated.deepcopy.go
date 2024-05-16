//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionInitParameters) DeepCopyInto(out *ConditionInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionInitParameters.
func (in *ConditionInitParameters) DeepCopy() *ConditionInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetIAMMember) DeepCopyInto(out *DatasetIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetIAMMember.
func (in *DatasetIAMMember) DeepCopy() *DatasetIAMMember {
	if in == nil {
		return nil
	}
	out := new(DatasetIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasetIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetIAMMemberInitParameters) DeepCopyInto(out *DatasetIAMMemberInitParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ConditionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DatasetID != nil {
		in, out := &in.DatasetID, &out.DatasetID
		*out = new(string)
		**out = **in
	}
	if in.DatasetIDRef != nil {
		in, out := &in.DatasetIDRef, &out.DatasetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatasetIDSelector != nil {
		in, out := &in.DatasetIDSelector, &out.DatasetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetIAMMemberInitParameters.
func (in *DatasetIAMMemberInitParameters) DeepCopy() *DatasetIAMMemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(DatasetIAMMemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetIAMMemberList) DeepCopyInto(out *DatasetIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatasetIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetIAMMemberList.
func (in *DatasetIAMMemberList) DeepCopy() *DatasetIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(DatasetIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasetIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetIAMMemberObservation) DeepCopyInto(out *DatasetIAMMemberObservation) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ConditionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.DatasetID != nil {
		in, out := &in.DatasetID, &out.DatasetID
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetIAMMemberObservation.
func (in *DatasetIAMMemberObservation) DeepCopy() *DatasetIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(DatasetIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetIAMMemberParameters) DeepCopyInto(out *DatasetIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ConditionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DatasetID != nil {
		in, out := &in.DatasetID, &out.DatasetID
		*out = new(string)
		**out = **in
	}
	if in.DatasetIDRef != nil {
		in, out := &in.DatasetIDRef, &out.DatasetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatasetIDSelector != nil {
		in, out := &in.DatasetIDSelector, &out.DatasetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetIAMMemberParameters.
func (in *DatasetIAMMemberParameters) DeepCopy() *DatasetIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(DatasetIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetIAMMemberSpec) DeepCopyInto(out *DatasetIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetIAMMemberSpec.
func (in *DatasetIAMMemberSpec) DeepCopy() *DatasetIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(DatasetIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetIAMMemberStatus) DeepCopyInto(out *DatasetIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetIAMMemberStatus.
func (in *DatasetIAMMemberStatus) DeepCopy() *DatasetIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(DatasetIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}