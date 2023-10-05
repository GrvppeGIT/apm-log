package tracer

import "os"

func SetEnvAPM(serviceName string, serverUrl string, secretToken string) bool {

	if err := os.Setenv(
		"ELASTIC_APM_SERVICE_NAME", serviceName); err != nil {
		return false
	}

	if err := os.Setenv(
		"ELASTIC_APM_SERVER_URL", serverUrl); err != nil {
		return false
	}

	if err := os.Setenv(
		"ELASTIC_APM_SECRET_TOKEN", secretToken); err != nil {
		return false
	}

	return true
}
