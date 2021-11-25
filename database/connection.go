package database

import (
	"fmt"
	"log"
	"strconv"

	"golang/config"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var db *gorm.DB

func Init() {
	c := &config.DatabaseConfig{}

	c.GetDatabaseConfig()

	if c.Driver == "mysql" {
		dsn := c.GetMySqlDSN()
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Printf("Cannot connect to %s database", c.Driver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("Database connected succesfully")
		}
	}
}

func Paginate(c echo.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetConnection() *gorm.DB {
	return db
}
