package env

import (
	"errors"

	"github.com/joho/godotenv"
)

var (
	// Error generated when loading the environment variable definition file.
	ErrLoadingDotEnv = errors.New("error on loading .env")
)

// Config loads the environment variables from the .env file.
func Config() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return ErrLoadingDotEnv
	}
	return nil
}
