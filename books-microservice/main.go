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
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title           User micro-service
// @version         1.0
// @description     Test
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:1322
// @BasePath  /
func main() {
	fmt.Println("Starting")

	conf := config.Load()
	newContainer := container.NewContainer(conf.GetEnv(), conf.GetURI(), conf.GetDBName())
	if newContainer.GetEnv() != "develop" {
		go kafkaDocker.Consume(context.Background(), newContainer)
	}
	e := echo.New()
	//e.Use(middleware.CORS())
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
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1322"))

}
