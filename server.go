package Shop

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.httpServer.Handler.ServeHTTP(w, r)
// }

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
