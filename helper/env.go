package helper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GoDotEnvVar(key string) string {
	err := godotenv.Load(filepath.Join(".env"))

	env := os.Getenv("APP_ENV")

	if env == "development" {
		if err != nil {
			fmt.Println(err)
		}
	}

	return os.Getenv(key)
}
