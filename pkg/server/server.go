package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Server
type Server struct {
	*gin.Engine
	Server   *http.Server
	config   *Config
	listener net.Listener
}

func newServer(config *Config) *Server {
	listener, err := net.Listen("tcp", config.Address())
	if err != nil {
		panic(fmt.Sprintf("failed to start server since: %v", err))
	}
	config.Port = listener.Addr().(*net.TCPAddr).Port
	gin.SetMode(config.Mode)
	return &Server{
		Engine:   gin.New(),
		config:   config,
		listener: listener,
	}
}

func (s *Server) Serve() error {
	// log on terminal
	s.Engine.Use(gin.Logger())
	s.Engine.Use(gin.Recovery(), gin.ErrorLogger())
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	for _, route := range s.Engine.Routes() {
		fmt.Printf("add route: Method(%s), path(%s)\n", route.Method, route.Path)
	}
	s.Server = &http.Server{
		Addr:    s.config.Address(),
		Handler: s,
	}
	err := s.Server.Serve(s.listener)
	if err == http.ErrServerClosed {
		fmt.Printf("close server %s", s.config.Address())
	}
	return err
}

func (s *Server) Stop() error {
	return s.Server.Close()
}

func (s *Server) GracefulStop(c context.Context) error {
	return s.Server.Shutdown(c)
}
