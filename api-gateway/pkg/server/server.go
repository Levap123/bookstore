package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run(router http.Handler) error {
	s.srv = &http.Server{
		ReadTimeout:    time.Second * 2,
		WriteTimeout:   time.Second * 2,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		Handler:        router,
	}
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
