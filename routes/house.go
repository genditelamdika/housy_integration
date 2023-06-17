package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func HouseRoutes(e *echo.Group) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerHouse(houseRepository)

	e.GET("/houses", h.FindHouses)
	e.GET("/house/:id", h.GetHouse)
	e.POST("/house", middleware.UploadFile(h.CreateHouse))
	e.PATCH("/house/:id", h.UpdateHouse)
	e.DELETE("/house/:id", h.DeleteHouse)
}
