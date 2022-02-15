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

	a := config.Load()

	newContainer := container.NewContainer(a.GetEnv(), a.GetURI(), a.GetDBName())
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.Init(e, newContainer)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, a.GetEnv())
	})
	e.Logger.Fatal(e.Start(":1323"))

}
