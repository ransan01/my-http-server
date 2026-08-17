package main

import (
	"fmt"
	"github.com/ransan01/my-http-server/internal/config"
	"github.com/ransan01/my-http-server/internal/router/netHttp"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
		return
	}
	router := netHttp.NewRouter()
	server := &Server{
		config: config,
		router: router,
	}
	server.Start()
}
