/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// grapisParameters are the configurable fields of a grapis.
type grapisParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// grapisObservation are the observable fields of a grapis.
type grapisObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A grapisSpec defines the desired state of a grapis.
type grapisSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       grapisParameters `json:"forProvider"`
}

// A grapisStatus represents the observed state of a grapis.
type grapisStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          grapisObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A grapis is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grpl}
type grapis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   grapisSpec   `json:"spec"`
	Status grapisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// grapisList contains a list of grapis
type grapisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []grapis `json:"items"`
}

// grapis type metadata.
var (
	grapisKind             = reflect.TypeOf(grapis{}).Name()
	grapisGroupKind        = schema.GroupKind{Group: Group, Kind: grapisKind}.String()
	grapisKindAPIVersion   = grapisKind + "." + SchemeGroupVersion.String()
	grapisGroupVersionKind = SchemeGroupVersion.WithKind(grapisKind)
)

func init() {
	SchemeBuilder.Register(&grapis{}, &grapisList{})
}
