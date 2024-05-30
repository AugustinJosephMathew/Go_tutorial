package config

import "os"

func Getpath() string {
	return os.Getenv("DB URL")
}
