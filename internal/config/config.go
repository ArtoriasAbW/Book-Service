package config

import "os"

func GetPostgresUser() string {
	return os.Getenv("POSTGRES_USER")
}

func GetPostgresPassword() string {
	return os.Getenv("POSTGRES_PASSWORD")
}
