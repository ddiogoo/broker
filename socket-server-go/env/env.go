package env

import (
	"errors"

	"github.com/joho/godotenv"
)

var (
	// Error generated when loading the environment variable definition file.
	ErrLoadingDotEnv = errors.New("error on loading .env")
)

// Method responsible for loading environment variables according to a configuration file.
// This method will only be called once.
func Config() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return ErrLoadingDotEnv
	}
	return nil
}
