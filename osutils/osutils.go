package osutils

import "os"

// Same as LookupEnv, but if no value exists, returns given default value.
// Bool true - if value comes from env, else false
func GetEnv(envKey string, defaultValue string) (string, bool) {
	val, exists := os.LookupEnv(envKey)
	if !exists {
		return defaultValue, false
	} else {
		return val, true
	}
}
