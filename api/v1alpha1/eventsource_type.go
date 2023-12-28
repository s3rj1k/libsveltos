/*
Copyright 2023. projectsveltos.io. All rights reserved.

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

const (
	// EventSourceFinalizer allows EventSourceReconciler to clean up resources associated with
	// EventSource before removing it from the apiserver.
	EventSourceFinalizer = "eventsource.finalizer.projectsveltos.io"

	EventSourceKind = "EventSource"
)

// ResourceSelector defines what resources are a match
type ResourceSelector struct {
	// Group of the resource deployed in the Cluster.
	Group string `json:"group"`

	// Version of the resource deployed in the Cluster.
	Version string `json:"version"`

	// Kind of the resource deployed in the Cluster.
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind"`

	// LabelFilters allows to filter resources based on current labels.
	// +optional
	LabelFilters []LabelFilter `json:"labelFilters,omitempty"`

	// Namespace of the resource deployed in the  Cluster.
	// Empty for resources scoped at cluster level.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Script contains a function "evaluate" in lua language.
	// The function will be passed one of the object selected based on
	// above criteria.
	// Must return struct with field "matching" representing whether
	// object is a match.
	// +optional
	Script string `json:"script,omitempty"`
}

// EventSourceSpec defines the desired state of EventSource
type EventSourceSpec struct {
	// ResourceSelectors identifies what resources to select
	ResourceSelectors []ResourceSelector `json:"resourceSelectors"`

	// This field is optional and can be used to specify a Lua function
	// that will be used to further select a subset of the resources that
	// have already been selected using the ResourceSelector field.
	// The function will receive the array of resources selected by ResourceSelectors.
	// If this field is not specified, all resources selected by the ResourceSelector
	// field will be considered.
	// This field allows to perform more complex filtering or selection operations
	// on the resources, looking at all resources together.
	// This can be useful for more sophisticated tasks, such as identifying resources
	// that are related to each other or that have similar properties.
	AggregatedSelection string `json:"aggregatedSelection,omitempty"`

	// CollectResources indicates whether matching resources need
	// to be collected and added to EventReport.
	// +kubebuilder:default:=false
	// +optional
	CollectResources bool `json:"collectResources,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:path=eventsources,scope=Cluster

// EventSource is the Schema for the EventSource API
type EventSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec EventSourceSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// EventSourceList contains a list of EventSource
type EventSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventSource{}, &EventSourceList{})
}
