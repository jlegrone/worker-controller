package v1alpha1

import (
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *TemporalWorker) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-temporal-io-temporal-io-v1alpha1-temporalworker,mutating=true,failurePolicy=fail,sideEffects=None,groups=temporal.io.temporal.io,resources=temporalworkers,verbs=create;update,versions=v1alpha1,name=mtemporalworker.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &TemporalWorker{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *TemporalWorker) Default() {
	if r.Spec.WorkerOptions.DeploymentSeries == "" {
		r.Spec.WorkerOptions.DeploymentSeries = fmt.Sprintf("%s.%s", r.GetName(), r.GetNamespace())
	}
}
