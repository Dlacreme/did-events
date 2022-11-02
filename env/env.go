package env

import (
	"log"
	"os"
)

var isDebug bool = false

func IsDebug() bool {
	return isDebug
}

func loadDebug() {
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		isDebug = true
	}
}

func LoadAndCheck() {
	loadDebug()

	requiredEnvVars := [...]string{
		"APP_PORT",
	}
	for _, envVarName := range requiredEnvVars {
		if os.Getenv(envVarName) == "" {
			log.Fatal("Missing following env variable: ", envVarName)
		}
	}
}
