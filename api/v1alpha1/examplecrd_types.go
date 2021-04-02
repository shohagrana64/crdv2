/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExampleCrdSpec defines the desired state of ExampleCrd
type ExampleCrdSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ExampleCrd. Edit ExampleCrd_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ExampleCrdStatus defines the observed state of ExampleCrd
type ExampleCrdStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ExampleCrd is the Schema for the examplecrds API
type ExampleCrd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExampleCrdSpec   `json:"spec,omitempty"`
	Status ExampleCrdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExampleCrdList contains a list of ExampleCrd
type ExampleCrdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExampleCrd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExampleCrd{}, &ExampleCrdList{})
}
