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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExecutorSpec defines the desired state of Executor
type ExecutorSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Types defines what types can be handled by executor e.g. "postman/collection", ":curl/command" etc
	Types []string `json:"types"`

	// ExecutorType one of "openapi" for rest based executors or "job" which will be default runners for kubtest soon
	ExecutorType string `json:"executor_type,omitempty"`
	// URI for rest based executors
	URI string `json:"uri,omitempty"`
	// Image for kube-job
	Image string `json:"image,omitempty"`
	// VolumeQuantity for kube-job PersistentVolume
	VolumeQuantity string `json:"volume_quantity,omitempty"`
	// VolumeMountPath - where should PV be monted inside job pod for e.g. artifacts
	VolumeMountPath string `json:"volume_mount_path,omitempty"`
}

type Runner struct {
}

// ExecutorStatus defines the observed state of Executor
type ExecutorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Executor is the Schema for the executors API
type Executor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExecutorSpec   `json:"spec,omitempty"`
	Status ExecutorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExecutorList contains a list of Executor
type ExecutorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Executor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Executor{}, &ExecutorList{})
}
