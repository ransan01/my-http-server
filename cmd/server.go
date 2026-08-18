package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ransan01/my-http-server/internal/config"
)

type Server struct {
	config *config.Config
	router *http.ServeMux
}

func (s *Server) Start() {
	address := s.config.Configurations[config.CONFIG_NAMES["Address"]].Types.Json.Value
	port := s.config.Configurations[config.CONFIG_NAMES["Port"]].Types.Json.Value
	certFile := s.config.Configurations[config.CONFIG_NAMES["CertFile"]].Types.Json.Value
	keyFile := s.config.Configurations[config.CONFIG_NAMES["KeyFile"]].Types.Json.Value

	fmt.Printf("Starting server on %s:%s\n", address, port)
	server := &http.Server{
		Addr:           address + ":" + port,
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
