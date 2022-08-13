package config

import "os"

func GetApiKey() string {
	return os.Getenv("TG_API_KEY")
}

func GetPostgresUser() string {
	return os.Getenv("POSTGRES_USER")
}

func GetPostgresPassword() string {
	return os.Getenv("POSTGRES_PASSWORD")
}
