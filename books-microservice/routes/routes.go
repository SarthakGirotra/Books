package router

import (
	"b/container"
	"b/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, container container.Container) {
	books := controllers.NewBookController(container)
	e.POST("/saveCSV", func(c echo.Context) error { return books.SaveFromCSV(c) })
	e.GET("/topBooks", func(c echo.Context) error { return books.TopBooks(c) })
	e.POST("/Like", func(c echo.Context) error { return books.Like(c) })
}
