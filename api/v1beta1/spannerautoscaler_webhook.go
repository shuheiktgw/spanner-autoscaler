/*
Copyright 2022.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var spannerautoscalerlog = logf.Log.WithName("spannerautoscaler-resource")

func (r *SpannerAutoscaler) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-spanner-mercari-com-v1beta1-spannerautoscaler,mutating=true,failurePolicy=fail,sideEffects=None,groups=spanner.mercari.com,resources=spannerautoscalers,verbs=create;update,versions=v1beta1,name=mspannerautoscaler.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &SpannerAutoscaler{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SpannerAutoscaler) Default() {
	spannerautoscalerlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

//+kubebuilder:webhook:path=/validate-spanner-mercari-com-v1beta1-spannerautoscaler,mutating=false,failurePolicy=fail,sideEffects=None,groups=spanner.mercari.com,resources=spannerautoscalers,verbs=create;update,versions=v1beta1,name=vspannerautoscaler.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &SpannerAutoscaler{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SpannerAutoscaler) ValidateCreate() error {
	spannerautoscalerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SpannerAutoscaler) ValidateUpdate(old runtime.Object) error {
	spannerautoscalerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SpannerAutoscaler) ValidateDelete() error {
	spannerautoscalerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
