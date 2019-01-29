package config

import "os"

const (
	FACILITY_KEY            = "facility"
	ENVIRONMENT_TYPE_KEY    = "environment"
	MESSAGE_KEY             = "message"
	HOST_NAME_KEY           = "hostName"
	HOST_NAME_ENV           = "HOST"
	LEVEL_KEY               = "level"
	TIMESTAMP_KEY           = "timestamp"
	CALLER_KEY              = "caller"
	SERVICE_NAME_ENV        = "AF_SERVICE_NAME"
	AF_ENVIRONMENT_TYPE_ENV = "AF_ENVIRONMENT_TYPE"
	AF_LOG_LEVEL_ENV        = "AF_LOG_LEVEL"
	BLANK_STRING            = ""
	DEV                     = "DEV"
	DEBUG                   = "debug"
)

func getEnv(envName, defaultValue string) string {
	if envVal := os.Getenv(envName); envVal != BLANK_STRING {
		return envVal
	}
	return defaultValue
}

//GetServiceName return a service name as string
func GetServiceName() string {
	return getEnv(SERVICE_NAME_ENV, DEV)
}

//GetEnvironmentType return environment type as string
func GetEnvironmentType() string {
	return getEnv(AF_ENVIRONMENT_TYPE_ENV, DEV)
}

func GetLogLevel() string {
	return getEnv(AF_LOG_LEVEL_ENV, DEBUG)
}

func GetHostName() string {
	host := getEnv(HOST_NAME_ENV, BLANK_STRING)
	if host != BLANK_STRING {
		return host
	}
	host, err := os.Hostname()
	if err != nil {
		return BLANK_STRING
	}
	return host
}
