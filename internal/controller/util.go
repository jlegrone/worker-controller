// Unless explicitly stated otherwise all files in this repository are licensed under the MIT License.
//
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2024 Datadog, Inc.

package controller

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"

	temporaliov1alpha1 "github.com/DataDog/temporal-worker-controller/api/v1alpha1"
	"github.com/DataDog/temporal-worker-controller/internal/controller/k8s.io/utils"
)

func computeVersionID(spec *temporaliov1alpha1.TemporalWorkerSpec) string {
	return spec.WorkerOptions.DeploymentName + "." + computeBuildID(spec)
}

func computeBuildID(spec *temporaliov1alpha1.TemporalWorkerSpec) string {
	return utils.ComputeHash(&spec.Template, nil)
}

func getTestWorkflowID(series, taskQueue, buildID string) string {
	return fmt.Sprintf("test-deploy:%s:%s:%s", series, taskQueue, buildID)
}

func newObjectRef(d *appsv1.Deployment) *v1.ObjectReference {
	if d == nil {
		return nil
	}
	return &v1.ObjectReference{
		Kind:            d.Kind,
		Namespace:       d.Namespace,
		Name:            d.Name,
		UID:             d.UID,
		APIVersion:      d.APIVersion,
		ResourceVersion: d.ResourceVersion,
	}
}
