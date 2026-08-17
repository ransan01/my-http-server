package main

import (
	"fmt"
	"github.com/ransan01/my-http-server/internal/config"
	"net/http"
	"time"
)

type Server struct {
	config *config.Config
	router *http.ServeMux
}

func (s *Server) Start() {
	fmt.Printf("Starting server on %s\n", s.config.Address+":"+s.config.Port)
	server := &http.Server{
		Addr:           s.config.Address + ":" + s.config.Port,
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServeTLS(s.config.CertFile, s.config.KeyFile); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	fmt.Printf("Server started on %s\n", s.config.Address+":"+s.config.Port)
}
