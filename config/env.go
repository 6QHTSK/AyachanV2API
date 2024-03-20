package config

import "os"

func envEnabled() bool {
	return os.Getenv("USE_ENV") == "true"
}

func readEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func initEnv() {
	Config = &YamlConfig{
		RunAddr: readEnv("RUN_ADDR", Config.RunAddr),
		Debug:   false,
		API: YamlConfigAPI{
			BestdoriProxy: readEnv("BD_API", Config.API.BestdoriProxy),
		},
	}
}
