package main

import (
	"fmt"
	"net"
	"src/api"
	"src/config"
	"src/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Inisialisasi Env
	err := godotenv.Load()
	if err != nil {
		e.Logger.Error(err)
	}

	// Inisialisasi DB
	db.Init(e, true, true)
	// Inisialisasi Server
	e = api.Init(e)

	// Server Listener
	port := config.GetConfig(e).Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
	e.Logger.Info("Port is:", e.Listener.Addr().(*net.TCPAddr).Port)
}
