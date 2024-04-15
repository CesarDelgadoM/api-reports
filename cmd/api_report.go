package main

import (
	"github.com/CesarDelgadoM/api-reports/config"
	"github.com/CesarDelgadoM/api-reports/server"
)

func main() {
	// Config load
	loadcfg := config.LoadConfig("config-local.yml")
	cfg := config.ParseConfig(loadcfg)

	// Server
	server := server.NewServer(cfg)
	server.Run()
}
