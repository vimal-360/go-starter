package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/do/v2"
	"go-workflow-rnd/internal/handlers"
)

func SetupRoutes(injector do.Injector) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api/v1")

	userHandler := do.MustInvoke[*handlers.UserHandler](injector)

	users := api.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	return e
}
