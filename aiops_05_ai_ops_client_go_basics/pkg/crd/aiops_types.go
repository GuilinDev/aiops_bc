package crd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AIOpsSpec defines the desired state of AIOps
type AIOpsSpec struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Target string `json:"target"`
}

// AIOps is the Schema for the aiops API
type AIOps struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AIOpsSpec `json:"spec,omitempty"`
}

// AIOpsList contains a list of AIOps
type AIOpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIOps `json:"items"`
} 