package main

import (
	"log"
	"os"
	"server/v0/pkg/server"
	"server/v0/pkg/utils"
)

func main() {

	serverConfig, err := utils.LoadServerConfig()
	if err != nil {
		os.Exit(1)
	}
	server, err := server.NewServer(*serverConfig)
	if err != nil {
		log.Printf("Unable to start server: %v", err)
		os.Exit(1)
	}
	server.Start()
}
