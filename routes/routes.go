package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AuthRoutes(e)
	HouseRoutes(e)
	AmenitiesRoutes(e)
	TransactionRoutes(e)

}
