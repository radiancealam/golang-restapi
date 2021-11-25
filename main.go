package main

import (
	"golang/database"
	"golang/helper"
	"golang/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	env := helper.GoDotEnvVar("APP_ENV")

	// Echo instance
	e := routes.Init()

	// Connect to Database
	database.Init()
	gorm := database.GetConnection()

	if env == "development" {
		database.Migrate()
	}

	dbGorm, err_gorm := gorm.DB()
	if err_gorm != nil {
		panic(err_gorm)
	}

	dbGorm.Ping()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
