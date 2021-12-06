package util

import (
	"log"
	"os"
	"strings"

	"github.com/Jaytpa01/iGustus/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// LoadENV loads the .env file into Golang environment to be used with os.Getenv
func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Info("error loading .env file", zap.Error(err))
	}
}

// RequireVariables ensures that all the variables in the keys parameter are not empty, if any are, it will log them and then it will panic
func RequireVariables(keys []string) {
	quit := false
	for _, key := range keys {
		if !VerifyVariableIsSet(key) {
			log.Printf("Missing ENV Variable: `%s`", key)
			quit = true
		}
	}

	if quit {
		log.Fatalf("Required ENV Variables are missing, see above")
	}
}

// VerifyVariableIsSet verifies whether a variable is empty
func VerifyVariableIsSet(key string) bool {
	return os.Getenv(key) != ""
}

func IsDevEnvironment() bool {
	return strings.ToLower(os.Getenv("GO_ENV")) == "development"
}
