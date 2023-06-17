package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transaction/:id", h.GetTransaction)
	e.GET("/user/:id/transaction", h.FindTransactionByUser)
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	// e.PATCH("/transaction/:id", h.UpdateTransaction)
	// e.DELETE("/transaction/:id", h.DeleteTransaction)
	e.POST("/notification", h.Notification)
}
