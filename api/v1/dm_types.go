/*
Copyright 2021.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DmSpec defines the desired state of Dm
type DmSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Dm. Edit dm_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DmStatus defines the observed state of Dm
type DmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Dm is the Schema for the dms API
type Dm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmSpec   `json:"spec,omitempty"`
	Status DmStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DmList contains a list of Dm
type DmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dm{}, &DmList{})
}
