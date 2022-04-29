package common

import (
	"fmt"
	"os"
)

func Env(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variables not set, key = %s", value)
	}

	return value, nil
}
