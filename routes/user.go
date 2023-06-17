package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	// e.POST("/user", h.CreateUser)
	e.PATCH("/user/:id", middleware.UploadFile(h.UpdateUser))
	e.DELETE("/user/:id", h.DeleteUser)
	e.PATCH("/change-password", middleware.Auth(h.ChangePassword))
}
