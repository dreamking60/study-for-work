package server

import "net/http"

type Server struct {
	addr       string
	httpServer *http.Server
}

func New() *Server {
	return &Server{
		addr: ":8080",
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: http.NewServeMux(),
		},
	}
}

func (s *Server) Addr() string {
	return s.addr
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
