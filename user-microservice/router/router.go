package router

import (
	"t/middlewareLocal"

	"t/controller"

	"t/container"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, container container.Container) {

	middlewareLocal.Val(e)
	health := controller.NewHealthController(container)
	user := controller.NewUserController(container)
	e.GET("/health", func(c echo.Context) error { return health.GetHealth(c) })
	e.POST("/login", func(c echo.Context) error { return user.Login(c) })
	e.POST("/signup", func(c echo.Context) error { return user.Signup(c) })
	e.GET("/user/:id", func(c echo.Context) error { return user.ValidateUser(c) })
}
