package main

import (
	"fmt"
	"loundryapp/config"
	"loundryapp/factory"
	"loundryapp/infrastructure/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()

	factory.InitFactory(e, db)

	fmt.Println("Starting programm ...")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
