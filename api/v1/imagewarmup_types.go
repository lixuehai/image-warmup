/*
Copyright 2025.

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

// ImageWarmupSpec defines the desired state of ImageWarmup.
type ImageWarmupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Images []ImageSpec `json:"images,omitempty"`

	// Schedule to use.
	// +optional
	Schedule string `json:"schedule,omitempty"`

	// Concurrency to use.
	// +optional
	// +kubebuilder:default:=3
	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:validation:Maximum:=10
	Concurrency int `json:"concurrency,omitempty"`

	// Timeout to use.
	// +optional
	// +kubebuilder:default:=300
	Timeout int `json:"timeout,omitempty"`
}

// ImageSpec defines the image to warmup.
type ImageSpec struct {
	// The image to warmup.
	Image string `json:"image"`

	// The pull secret to use.
	// +optional
	PullSecret string `json:"pullSecret,omitempty"`

	// The nodes to warmup the image on.
	// +optional
	Nodes []string `json:"nodes,omitempty"`

	// The retry count to use.
	// +optional
	// +kubebuilder:default:=3
	RetryCount int `json:"retryCount,omitempty"`
}

// ImageWarmupStatus defines the observed state of ImageWarmup.
type ImageWarmupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Phase is the current phase of the ImageWarmup.
	Phase string `json:"phase,omitempty"`

	// WarmedUpImages is the list of images that have been warmed up.
	WarmedUpImages []WarmedImage `json:"warmedUpImages,omitempty"`

	// FailedImages is the list of images that have failed to warm up.
	FailedImages []FailedImage `json:"failedImages,omitempty"`

	// Message is the message of the ImageWarmup.
	Message string `json:"message,omitempty"`

	// Conditions is the list of conditions of the ImageWarmup.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type FailedImage struct {
	// Image is the image that failed to warm up.
	Image string `json:"image"`

	// Nodes is the list of nodes that the image failed to warm up on.
	Nodes []string `json:"nodes,omitempty"`

	// Error is the error message of the image failed to warm up.
	Error string `json:"error,omitempty"`

	// RetryCount is the retry count of the image failed to warm up.
	RetryCount int `json:"retryCount,omitempty"`

	// LastAttemptTime is the time when the image failed to warm up.
	LastAttemptTime metav1.Time `json:"lastAttemptTime,omitempty"`
}

// WarmedImage is the image that has been warmed up.
type WarmedImage struct {
	// Image is the image that was warmed up.
	Image string `json:"image"`

	// Nodes is the list of nodes that the image was warmed up on.
	Nodes []string `json:"nodes,omitempty"`

	// WarmupTime is the time when the image was warmed up.
	WarmupTime metav1.Time `json:"warmupTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="The phase of the ImageWarmup"
// +kubebuilder:printcolumn:name="WarmedUpImages",type="string",JSONPath=".status.warmedUpImages",description="The images that have been warmed up"
// +kubebuilder:printcolumn:name="FailedImages",type="string",JSONPath=".status.failedImages",description="The images that have failed to warm up"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message",description="The message of the ImageWarmup"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the ImageWarmup"

// ImageWarmup is the Schema for the imagewarmups API.
type ImageWarmup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageWarmupSpec   `json:"spec,omitempty"`
	Status ImageWarmupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageWarmupList contains a list of ImageWarmup.
type ImageWarmupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageWarmup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageWarmup{}, &ImageWarmupList{})
}
