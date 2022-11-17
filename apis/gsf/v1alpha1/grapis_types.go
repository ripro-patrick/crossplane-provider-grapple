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

// GrapisParameters are the configurable fields of a Grapis.
type GrapisParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// GrapisObservation are the observable fields of a Grapis.
type GrapisObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A GrapisSpec defines the desired state of a Grapis.
type GrapisSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GrapisParameters `json:"forProvider"`
}

// A GrapisStatus represents the observed state of a Grapis.
type GrapisStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GrapisObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Grapis is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grpl}
type Grapis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrapisSpec   `json:"spec"`
	Status GrapisStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrapisList contains a list of Grapis
type GrapisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grapis `json:"items"`
}

// Grapis type metadata.
var (
	GrapisKind             = reflect.TypeOf(Grapis{}).Name()
	GrapisGroupKind        = schema.GroupKind{Group: Group, Kind: GrapisKind}.String()
	GrapisKindAPIVersion   = GrapisKind + "." + SchemeGroupVersion.String()
	GrapisGroupVersionKind = SchemeGroupVersion.WithKind(GrapisKind)
)

func init() {
	SchemeBuilder.Register(&Grapis{}, &GrapisList{})
}
