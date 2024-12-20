package main

import "github.com/mcLyu/tcp-proof-of-work/internal/tcpclient"

func main() {
	client := tcpclient.New()
	client.Connect("8080")
}
