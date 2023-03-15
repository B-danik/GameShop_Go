package todo

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, router *httprouter.Router) error {
	s.httpServer = &http.Server{
		Handler:        router,
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
