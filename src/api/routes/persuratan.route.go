package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Persuratan(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/persuratan", controllers.CreatePersuratan)
	auth.PUT("/persuratan/selesai/:id", controllers.SelesaiPersuratanById)
	auth.PUT("/persuratan/tolak/:id", controllers.TolakPersuratanById)
	auth.GET("/persuratan", controllers.GetAllPersuratan)
	auth.GET("/persuratan/:id", controllers.GetPersuratanByID)
	auth.PUT("/persuratan/:id", controllers.UpdatePersuratanById)
	auth.DELETE("/persuratan/:id", controllers.SoftDeletePersuratanById)

	return e
}
