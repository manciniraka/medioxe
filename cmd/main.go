package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	commonDB "github.com/manciniraka/go-common/database"
	commonValidator "github.com/manciniraka/go-common/validator"
	"github.com/manciniraka/medioxe/internal/router"
)

func main() {
	db, err := commonDB.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}

	_ = db

	e := echo.New()

	e.Validator = commonValidator.New()

	router.InitRouter(e)

	port := os.Getenv("APP_PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
