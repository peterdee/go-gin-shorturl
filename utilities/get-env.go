package utilities

import "os"

func GetEnv(name string, defaultValue ...string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}
