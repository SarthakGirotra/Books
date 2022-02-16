package main

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"t/container"
	"t/router"

	"t/config"

	"github.com/labstack/echo/v4"
)

func main() {

	conf := config.Load()
	newContainer := container.NewContainer(conf.GetEnv(), conf.GetURI(), conf.GetDBName())

	e := echo.New()
	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Init(e, newContainer)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, conf.GetEnv())
	})

	e.Static("/swagger-ui.css", "dist/swagger-ui.css")
	e.Static("/swagger-ui-bundle.js", "dist/swagger-ui-bundle.js")
	e.Static("/swagger-ui-standalone-preset.js", "dist/swagger-ui-standalone-preset.js")
	e.Static("/swagger.json", "./swagger.json")
	e.Static("/swaggerui", "dist/index.html")

	e.Logger.Fatal(e.Start(":1323"))

}
