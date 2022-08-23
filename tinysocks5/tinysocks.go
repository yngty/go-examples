package main

var Version string = "dev"

var needsTLS = map[string]bool{
	"https": true,
	"wss":   true,
}
