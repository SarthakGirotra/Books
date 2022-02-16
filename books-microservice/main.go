package main

import (
	"b/config"
	"b/container"
	"b/kafkaDocker"
	router "b/routes"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Starting")

	conf := config.Load()
	newContainer := container.NewContainer(conf.GetEnv(), conf.GetURI(), conf.GetDBName())
	if newContainer.GetEnv() != "develop" {
		go kafkaDocker.Consume(context.Background(), newContainer)
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	router.Init(e, newContainer)

	e.GET("/", func(c echo.Context) error {
		return c.File("./index.html")
	})
	e.Static("/swagger-ui.css", "dist/swagger-ui.css")
	e.Static("/swagger-ui-bundle.js", "dist/swagger-ui-bundle.js")
	e.Static("/swagger-ui-standalone-preset.js", "dist/swagger-ui-standalone-preset.js")
	e.Static("/swagger.json", "./swagger.json")
	e.Static("/swaggerui", "dist/index.html")
	e.Logger.Fatal(e.Start(":1322"))

}
