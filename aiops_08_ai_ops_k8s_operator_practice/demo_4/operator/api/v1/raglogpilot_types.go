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

// RagLogPilotSpec defines the desired state of RagLogPilot
type RagLogPilotSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	WorkloadNameSpace string `json:"workloadNameSpace"`
	RagFlowEndpoint   string `json:"ragFlowEndpoint"`
	RagFlowToken      string `json:"ragFlowToken"`
	LLMType          string `json:"llmType,omitempty"`
	LLMEndpoint      string `json:"llmEndpoint,omitempty"`
	LLMModel         string `json:"llmModel,omitempty"`
}

// RagLogPilotStatus defines the observed state of RagLogPilot
type RagLogPilotStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ConversationId string `json:"conversationId"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RagLogPilot is the Schema for the raglogpilots API
type RagLogPilot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RagLogPilotSpec   `json:"spec,omitempty"`
	Status RagLogPilotStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RagLogPilotList contains a list of RagLogPilot
type RagLogPilotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RagLogPilot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RagLogPilot{}, &RagLogPilotList{})
}
