/*
Copyright 2024.

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

// VirtualServerManagerSpec defines the desired state of VirtualServerManager
type VirtualServerManagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// VirtualServerName is the name of the virtual server to manage
	VirtualServerName string `json:"virtualServerName"`
	// Namespace is the namespace of the virtual server to manage
	Namespace string `json:"namespace"`
	// VariableKey is the key of the variable to manage
	VariableKey string `json:"variableKey"`
	// VariableValue is the value of the variable to manage
	VariableValue string `json:"variableValue"`
}

// VirtualServerManagerStatus defines the observed state of VirtualServerManager
type VirtualServerManagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Updated is true if the variable has been updated
	Updated bool `json:"updated"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VirtualServerManager is the Schema for the virtualservermanagers API
type VirtualServerManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualServerManagerSpec   `json:"spec,omitempty"`
	Status VirtualServerManagerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VirtualServerManagerList contains a list of VirtualServerManager
type VirtualServerManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualServerManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualServerManager{}, &VirtualServerManagerList{})
}
