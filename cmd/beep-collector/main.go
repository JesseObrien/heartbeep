package main

import (
	"flag"

	"github.com/jesseobrien/heartbeep/internal/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8888", "A port number. e.g. 9000")
	flag.Parse()
}

func main() {
	beepCollector := server.CollectorServer{}

	beepCollector.Run(port)
}
