package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func AmenitiesRoutes(e *echo.Group) {
	amenitiesRepository := repositories.RepositoryAmenities(mysql.DB)
	h := handlers.HandlerAmenities(amenitiesRepository)

	e.GET("/amenities", h.FindAmenities)
	e.GET("/amenitie/:id", h.GetAmenitie)
	e.POST("/amenitie", h.CreateAmenitie)
	e.PATCH("/amenitie/:id", h.UpdateAmenitie)
	e.DELETE("/amenitie/:id", h.DeleteAmenitie)
}
