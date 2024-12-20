package main

import (
	"github.com/mcLyu/tcp-proof-of-work/internal/tcpserver"
)

const (
	DefaultPort = "8080"
)

func main() {
	server := tcpserver.New()
	server.Listen(DefaultPort)
}
