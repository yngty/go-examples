package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

func main() {
	var configPath string
	var showVersion bool

	flag.StringVar(&configPath, "c", "", "configuration file, default to 'config.toml'")
	flag.BoolVar(&showVersion, "v", false, "show version information")
	flag.Parse()

	if showVersion {
		fmt.Println("tinysocks", Version)
		return
	}

	if configPath == "" {
		configPath = "config.toml"
		log.Printf("Using default configuration 'config.toml'")
	}

	config, err := toml.LoadFile(configPath)
	if err != nil {
		log.Fatalf("Load configuration failed: %s", err)
	}
	if c, ok := config.Get("client").(*toml.Tree); ok {

	} else if s, ok := config.Get("server").(*toml.Tree); ok {
	} else {
		log.Fatal("No valid configuration '[client]' or '[server]'")
	}

}
