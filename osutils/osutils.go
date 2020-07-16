package osutils

import "os"

func GetEnv(envKey string, defaultValue string) (string, bool) {
	val, exists := os.LookupEnv(envKey)
	if !exists {
		return defaultValue, false
	} else {
		return val, true
	}
}
