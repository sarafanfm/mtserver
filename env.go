package mtserver

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const ENV_FILE_VAR = "ENV_FILE"

func RequiredEnv(envVar string) string {
	val, ok := os.LookupEnv(envVar)
	if !ok {
		log.Fatalf("cannot resolve env var %s", envVar)
	}
	return val
}

func preloadEnv() {
	customEnvFile, ok := os.LookupEnv(ENV_FILE_VAR)
	if ok && customEnvFile != "" {
		files := strings.Split(customEnvFile, ",")
		for _, v := range files {
			_ = godotenv.Load(v)
		}
	}
}
