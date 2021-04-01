/*
Copyright 2021.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Dm8Spec defines the desired state of Dm8
type Dm8Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Dm8. Edit dm8_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// Dm8Status defines the observed state of Dm8
type Dm8Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Dm8 is the Schema for the dm8s API
type Dm8 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Dm8Spec   `json:"spec,omitempty"`
	Status Dm8Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Dm8List contains a list of Dm8
type Dm8List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dm8 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dm8{}, &Dm8List{})
}
