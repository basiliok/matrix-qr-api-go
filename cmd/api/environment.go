package main

import "os"

func getEnvironmentVariable(variableName string, defaultValue string) string {
	value, isDefined := os.LookupEnv(variableName)
	if !isDefined || value == "" {
		return defaultValue
	}

	return value
}
