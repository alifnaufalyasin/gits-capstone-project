package main

import (
	"fmt"
	"net"
	"os"
	"src/api"
	"src/config"
	"src/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Inisialisasi Logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	e := echo.New()
	// Inisialisasi Env
	err := godotenv.Load()
	if err != nil {
		log.Error().Msgf("%v", err)
	}

	// Inisialisasi DB
	db.Init(e, true, true)
	// Inisialisasi Server
	e = api.Init(e)

	// Server Listener
	port := config.GetConfig().Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
	e.Logger.Info("Port is:", e.Listener.Addr().(*net.TCPAddr).Port)
}
