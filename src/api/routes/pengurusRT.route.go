package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func PengurusRT(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("/pengurus")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.GET("", controllers.GetAllPengurusRT)
	auth.GET("/me", controllers.GetPengurusByID)
	auth.PUT("/changepassword", controllers.GantiPasswordPengurus)
	auth.GET("/detail/:id", controllers.GetPengurusByID)
	auth.PUT("/:id", controllers.UpdatePengurusById)
	auth.DELETE("/:id", controllers.SoftDeletePengurusById)

	e.POST("/pengurus", controllers.CreatePengurus)
	e.POST("/pengurus/login", controllers.LoginPengurus)
	e.POST("/pengurus/forgetpassword", controllers.ForgetPasswordPengurus)
	e.POST("/pengurus/resetpasswordbykode", controllers.ResetPasswordPengurusByKode)

	return e
}
