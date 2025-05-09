package main

import (
	"www.github/M1ralai/tcp/cmd/server"
)

func main() {
	t := server.NewTCPServer(":3000")
	t.ListenAndAccept()
	select {}
}
