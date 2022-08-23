package server

import (
	"crypto/tls"
	"net"
)

type Server struct {
	Config    *Config
	TLSConfig *tls.Config
}

func New(protocol, addr string) *Server {
	return &Server{
		Config: &Config{
			Protocol: protocol,
			Addr:     addr,
		},
	}
}

var protocol2Handler = map[string]func(*Server, net.Conn){
	"https": (*Server).
}

type Config struct {
	Protocol   string
	Addr       string
	Verify     func(string, string) bool
	HTTPPath   string
	WSPath     string
	WSCompress bool
}
