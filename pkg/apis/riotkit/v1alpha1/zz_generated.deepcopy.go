//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Riotkit.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupRefSpec) DeepCopyInto(out *BackupRefSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupRefSpec.
func (in *BackupRefSpec) DeepCopy() *BackupRefSpec {
	if in == nil {
		return nil
	}
	out := new(BackupRefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBackupProcedureTemplate) DeepCopyInto(out *ClusterBackupProcedureTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBackupProcedureTemplate.
func (in *ClusterBackupProcedureTemplate) DeepCopy() *ClusterBackupProcedureTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterBackupProcedureTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBackupProcedureTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBackupProcedureTemplateList) DeepCopyInto(out *ClusterBackupProcedureTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBackupProcedureTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBackupProcedureTemplateList.
func (in *ClusterBackupProcedureTemplateList) DeepCopy() *ClusterBackupProcedureTemplateList {
	if in == nil {
		return nil
	}
	out := new(ClusterBackupProcedureTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBackupProcedureTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBackupProcedureTemplateSpec) DeepCopyInto(out *ClusterBackupProcedureTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBackupProcedureTemplateSpec.
func (in *ClusterBackupProcedureTemplateSpec) DeepCopy() *ClusterBackupProcedureTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBackupProcedureTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBackupProcedureTemplateStatus) DeepCopyInto(out *ClusterBackupProcedureTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBackupProcedureTemplateStatus.
func (in *ClusterBackupProcedureTemplateStatus) DeepCopy() *ClusterBackupProcedureTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBackupProcedureTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPGKeySecretSpec) DeepCopyInto(out *GPGKeySecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPGKeySecretSpec.
func (in *GPGKeySecretSpec) DeepCopy() *GPGKeySecretSpec {
	if in == nil {
		return nil
	}
	out := new(GPGKeySecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoredBackup) DeepCopyInto(out *RestoredBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoredBackup.
func (in *RestoredBackup) DeepCopy() *RestoredBackup {
	if in == nil {
		return nil
	}
	out := new(RestoredBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestoredBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoredBackupList) DeepCopyInto(out *RestoredBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RestoredBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoredBackupList.
func (in *RestoredBackupList) DeepCopy() *RestoredBackupList {
	if in == nil {
		return nil
	}
	out := new(RestoredBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestoredBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoredBackupSpec) DeepCopyInto(out *RestoredBackupSpec) {
	*out = *in
	out.ScheduledBackupRef = in.ScheduledBackupRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoredBackupSpec.
func (in *RestoredBackupSpec) DeepCopy() *RestoredBackupSpec {
	if in == nil {
		return nil
	}
	out := new(RestoredBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoredBackupStatus) DeepCopyInto(out *RestoredBackupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoredBackupStatus.
func (in *RestoredBackupStatus) DeepCopy() *RestoredBackupStatus {
	if in == nil {
		return nil
	}
	out := new(RestoredBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackup) DeepCopyInto(out *ScheduledBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackup.
func (in *ScheduledBackup) DeepCopy() *ScheduledBackup {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackupList) DeepCopyInto(out *ScheduledBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackupList.
func (in *ScheduledBackupList) DeepCopy() *ScheduledBackupList {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackupSpec) DeepCopyInto(out *ScheduledBackupSpec) {
	*out = *in
	out.TemplateRef = in.TemplateRef
	out.GPGKeySecretRef = in.GPGKeySecretRef
	out.TokenSecretRef = in.TokenSecretRef
	in.VarsSecretRef.DeepCopyInto(&out.VarsSecretRef)
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(VarsSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackupSpec.
func (in *ScheduledBackupSpec) DeepCopy() *ScheduledBackupSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackupStatus) DeepCopyInto(out *ScheduledBackupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackupStatus.
func (in *ScheduledBackupStatus) DeepCopy() *ScheduledBackupStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSecretSpec) DeepCopyInto(out *TokenSecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSecretSpec.
func (in *TokenSecretSpec) DeepCopy() *TokenSecretSpec {
	if in == nil {
		return nil
	}
	out := new(TokenSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsSecretSpec) DeepCopyInto(out *VarsSecretSpec) {
	*out = *in
	if in.ImportOnlyKeys != nil {
		in, out := &in.ImportOnlyKeys, &out.ImportOnlyKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSecretSpec.
func (in *VarsSecretSpec) DeepCopy() *VarsSecretSpec {
	if in == nil {
		return nil
	}
	out := new(VarsSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in VarsSpec) DeepCopyInto(out *VarsSpec) {
	{
		in := &in
		*out = make(VarsSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSpec.
func (in VarsSpec) DeepCopy() VarsSpec {
	if in == nil {
		return nil
	}
	out := new(VarsSpec)
	in.DeepCopyInto(out)
	return *out
}
