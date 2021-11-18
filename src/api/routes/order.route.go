package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
)

func Order(e *echo.Echo) *echo.Echo {
	e.POST("/order", controllers.CreateOrder)
	e.GET("/order", controllers.GetAllOrder)
	e.GET("/order/:id", controllers.GetOrderByID)
	e.PUT("/order/:id", controllers.UpdateOrderById)
	e.DELETE("/order/:id", controllers.SoftDeleteOrderById)

	return e
}
