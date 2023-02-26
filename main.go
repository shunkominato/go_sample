package main

import (
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			"127.0.0.1", 5432, "postgres", "rails_sample", "postgres",
		),
	)
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/", welcome())

	err := e.Start(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}