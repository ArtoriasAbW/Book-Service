package config

import "os"

func GetApiKey() string {
	return os.Getenv("TG_API_KEY")
}
