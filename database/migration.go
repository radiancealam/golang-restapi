package database

import (
	"fmt"
	"golang/helper"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	user := helper.GoDotEnvVar("DB_USER")
	password := helper.GoDotEnvVar("DB_PASSWORD")
	host := helper.GoDotEnvVar("DB_HOST")
	port := helper.GoDotEnvVar("DB_PORT")
	name := helper.GoDotEnvVar("DB_NAME")

	url := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)

	m, err := migrate.New(
		"file://database/migrations",
		url,
	)

	if err != nil {
		fmt.Println(err)
	}

	if err := m.Up(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database is migrated succesfully")
	}
}
