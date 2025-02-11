package util

import (
	"net/http"
	"time"

	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func NewVersionedWorker(opts worker.Options) (w worker.Worker, stopFunc func()) {
	go func() {
		// Delay pod readiness by 5 seconds
		time.Sleep(5 * time.Second)
		if err := http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})); err != nil {
			panic(err)
		}
	}()

	opts.BuildID = mustGetEnv("WORKER_BUILD_ID")
	opts.UseBuildIDForVersioning = true
	opts.DeploymentOptions = worker.DeploymentOptions{
		DeploymentSeriesName:      mustGetEnv("TEMPORAL_DEPLOYMENT_SERIES"),
		DefaultVersioningBehavior: workflow.VersioningBehaviorPinned,
	}

	c, stopClient := NewClient(opts.BuildID)

	w = worker.New(c, temporalTaskQueue, opts)

	return w, func() {
		w.Stop()
		stopClient()
	}
}
