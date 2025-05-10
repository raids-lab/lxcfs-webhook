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
	"context"
	"fmt"
	"slices"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// nolint:unused
// log is for logging in this package.
var podlog = logf.Log.WithName("pod-resource")

// SetupPodWebhookWithManager registers the webhook for Pod in the manager.
func SetupPodWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&corev1.Pod{}).
		WithValidator(&PodLxcfsValidator{}).
		WithDefaulter(&PodLxcfsDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// nolint:lll
// +kubebuilder:webhook:path=/mutate--v1-pod,mutating=true,failurePolicy=ignore,sideEffects=None,groups="",resources=pods,verbs=create,versions=v1,name=mpod-v1.kb.io,admissionReviewVersions=v1

// PodLxcfsDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Pod when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type PodLxcfsDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &PodLxcfsDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Pod.
func (d *PodLxcfsDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	pod, ok := obj.(*corev1.Pod)

	if !ok {
		return fmt.Errorf("expected an Pod object but got %T", obj)
	}
	podlog.Info("Defaulting for Pod", "name", pod.GetName(), "namespace", pod.GetNamespace())

	// Check if the Pod should be mutated
	if !mutationRequired(pod) {
		podlog.Info("Skipping mutation for Pod", "name", pod.GetName(), "namespace", pod.GetNamespace())
		return nil
	}

	// If the Pod is not mutated, we need to add the annotation
	if pod.Annotations == nil {
		pod.Annotations = make(map[string]string)
	}
	pod.Annotations[AdmissionWebhookAnnotationStatusKey] = StatusValueMutated

	// Add LXCFS VolumeMounts to all containers
	for i := range pod.Spec.Containers {
		container := &pod.Spec.Containers[i]
		if container.VolumeMounts == nil {
			container.VolumeMounts = make([]corev1.VolumeMount, 0)
		}
		container.VolumeMounts = append(container.VolumeMounts, VolumeMountsTemplate...)
	}

	// Add LXCFS VolumeMounts to pod
	if pod.Spec.Volumes == nil {
		pod.Spec.Volumes = make([]corev1.Volume, 0)
	}
	pod.Spec.Volumes = append(pod.Spec.Volumes, VolumesTemplate...)

	return nil
}

// mutationRequired determines whether the Pod should be mutated or not.
func mutationRequired(pod *corev1.Pod) bool {
	// Ignore system namespaces, such as kube-system, kube-public, and kube-node-lease
	if slices.Contains(IgnoredNamespaces, pod.Namespace) {
		podlog.Info("Skip mutation for system namespace", "namespace", pod.Namespace)
		return false
	}

	annotations := pod.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}

	// Check if the Pod is already mutated
	if status, ok := annotations[AdmissionWebhookAnnotationStatusKey]; ok &&
		status == StatusValueMutated {
		return false
	}

	// Check if the Pod should be mutated
	// If the annotation is not set, we assume it should be mutated
	// If the annotation is set to "false", we skip the mutation
	if mutate, ok := annotations[AdmissionWebhookAnnotationMutateKey]; ok {
		switch strings.ToLower(mutate) {
		case MutateValueFalse:
			return false
		}
	}

	return true
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// nolint:lll
// +kubebuilder:webhook:path=/validate--v1-pod,mutating=false,failurePolicy=ignore,sideEffects=None,groups="",resources=pods,verbs=create;update,versions=v1,name=vpod-v1.kb.io,admissionReviewVersions=v1

// PodLxcfsValidator struct is responsible for validating the Pod resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type PodLxcfsValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &PodLxcfsValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Pod.
func (v *PodLxcfsValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		return nil, fmt.Errorf("expected a Pod object but got %T", obj)
	}
	podlog.Info("Validation for Pod upon creation", "name", pod.GetName(), "namespace", pod.GetNamespace())

	return nil, validate(pod)
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Pod.
func (v *PodLxcfsValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (
	admission.Warnings, error,
) {
	pod, ok := newObj.(*corev1.Pod)
	if !ok {
		return nil, fmt.Errorf("expected a Pod object for the newObj but got %T", newObj)
	}
	podlog.Info("Validation for Pod upon update", "name", pod.GetName(), "namespace", pod.GetNamespace())

	return nil, validate(pod)
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Pod.
func (v *PodLxcfsValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	return nil, nil
}

// Validate checks the Pod annotations to ensure that the mutation annotation is set to a valid value.
// Only "true" or "false" are valid values.
func validate(pod *corev1.Pod) error {
	annotations := pod.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}

	if mutate, ok := annotations[AdmissionWebhookAnnotationMutateKey]; ok {
		switch strings.ToLower(mutate) {
		case MutateValueTrue, MutateValueFalse:
			return nil
		default:
			return fmt.Errorf("invalid value for %s annotation: %s", AdmissionWebhookAnnotationMutateKey, mutate)
		}
	}

	return nil
}
