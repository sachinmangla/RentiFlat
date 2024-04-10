package server

import (
	"net/http"
	"time"

	"github.com/sachinmangla/rentiflat/routes"
)

func RunServer(port string) error {
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        routes.GetRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server.ListenAndServe()
}
