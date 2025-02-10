package server

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Host string
	Port int
}

type Server struct {
	httpServer *http.Server
	Config
}

// Creates a new server instance.
//   - if Config.Host is empty, it will default to 0.0.0.0
//   - if Config.Port is the empty value, it will default to 3000
func New(cfg Config) (srv *Server) {
	if cfg.Host == "" {
		cfg.Host = "0.0.0.0"
	}
	if cfg.Port == 0 {
		cfg.Port = 3000
	}

	m := http.NewServeMux()
	m.HandleFunc("POST /json", jsonHandler())

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: m,
	}

	return &Server{
		httpServer: httpServer,
		Config:     cfg,
	}
}

func (s *Server) Start() {
	defer s.httpServer.Close()
	log.Println(
		"[INF]",
		"starting server",
		fmt.Sprintf("host=%s port=%d", s.Config.Host, s.Config.Port),
	)
	s.httpServer.ListenAndServe()
}
