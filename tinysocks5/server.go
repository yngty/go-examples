package main

import (
	"log"

	"github.com/pelletier/go-toml"
)

func launchServer(t *toml.Tree) {
	config := struct {
		Protocol string `toml:"protocol"`
		Addr     string `toml:"listen"`
		HTTP     struct {
			Path string `toml:"path" default:"/"`
		} `toml:"http"`

		WS struct {
			Path     string `toml:"path" default:"/"`
			Compress bool   `toml:"compress"`
		} `toml:"ws"`

		TLS struct {
			Cert string `toml:"cert"`
			Key  string `toml:"key"`
		} `toml:"tls"`
	}{}

	if err := t.Unmarshal(&config); err != nil {
		log.Fatalf("Parse '[server]' configuration failed: %s", err)
	}
	serv := s
}
