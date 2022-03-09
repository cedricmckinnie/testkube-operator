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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuite) DeepCopyInto(out *TestSuite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuite.
func (in *TestSuite) DeepCopy() *TestSuite {
	if in == nil {
		return nil
	}
	out := new(TestSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestSuite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteList) DeepCopyInto(out *TestSuiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestSuite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteList.
func (in *TestSuiteList) DeepCopy() *TestSuiteList {
	if in == nil {
		return nil
	}
	out := new(TestSuiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestSuiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteSpec) DeepCopyInto(out *TestSuiteSpec) {
	*out = *in
	if in.Before != nil {
		in, out := &in.Before, &out.Before
		*out = make([]TestSuiteStepSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]TestSuiteStepSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.After != nil {
		in, out := &in.After, &out.After
		*out = make([]TestSuiteStepSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteSpec.
func (in *TestSuiteSpec) DeepCopy() *TestSuiteSpec {
	if in == nil {
		return nil
	}
	out := new(TestSuiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteStatus) DeepCopyInto(out *TestSuiteStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteStatus.
func (in *TestSuiteStatus) DeepCopy() *TestSuiteStatus {
	if in == nil {
		return nil
	}
	out := new(TestSuiteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteStepDelay) DeepCopyInto(out *TestSuiteStepDelay) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteStepDelay.
func (in *TestSuiteStepDelay) DeepCopy() *TestSuiteStepDelay {
	if in == nil {
		return nil
	}
	out := new(TestSuiteStepDelay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteStepExecute) DeepCopyInto(out *TestSuiteStepExecute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteStepExecute.
func (in *TestSuiteStepExecute) DeepCopy() *TestSuiteStepExecute {
	if in == nil {
		return nil
	}
	out := new(TestSuiteStepExecute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteStepSpec) DeepCopyInto(out *TestSuiteStepSpec) {
	*out = *in
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = new(TestSuiteStepExecute)
		**out = **in
	}
	if in.Delay != nil {
		in, out := &in.Delay, &out.Delay
		*out = new(TestSuiteStepDelay)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteStepSpec.
func (in *TestSuiteStepSpec) DeepCopy() *TestSuiteStepSpec {
	if in == nil {
		return nil
	}
	out := new(TestSuiteStepSpec)
	in.DeepCopyInto(out)
	return out
}
