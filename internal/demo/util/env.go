package util

import (
	"fmt"
	"os"
)

var (
	temporalHostPort  = mustGetEnv("TEMPORAL_HOST_PORT")
	temporalNamespace = mustGetEnv("TEMPORAL_NAMESPACE")
	temporalTaskQueue = mustGetEnv("TEMPORAL_TASK_QUEUE")
	tlsKeyFilePath    = mustGetEnv("TEMPORAL_TLS_KEY_PATH")
	tlsCertFilePath   = mustGetEnv("TEMPORAL_TLS_CERT_PATH")
)

func mustGetEnv(key string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	panic(fmt.Sprintf("environment variable %q must be set", key))
}
