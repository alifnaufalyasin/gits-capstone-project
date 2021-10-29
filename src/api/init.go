package api

import (
	"src/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Init(e *echo.Echo) *echo.Echo {
	log.Info().Msg("menginisialisasikan server")

	e = routes.Init(e)

	log.Info().Msg("server terinisialisasi")

	return e
}
