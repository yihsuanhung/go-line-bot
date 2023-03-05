package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Config HTTP config
type Config struct {
	Host string
	Port int
	Mode string
}

// Default Config
func DefaultConfig() *Config {
	return &Config{
		Host: "0.0.0.0",
		Port: 8080,
		Mode: gin.ReleaseMode,
	}
}

// Build create server instance, then initialize it with necessary interceptor
func (config *Config) Build() *Server {
	s := newServer(config)
	return s
}

// Address
func (config *Config) Address() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}
