package server

import (
	"log"
	"net"

	"www.github/M1ralai/tcp/cmd/client"
	"www.github/M1ralai/tcp/cmd/menu"
)

func (t *TCPServer) handleCon(conn net.Conn) {
	log.Print("Connection established ", conn.RemoteAddr())
	u := menu.NoLoginMenu(conn)
	if u == nil {
		return
	}
	c := client.NewClient(*u, conn)
	t.Clients = append(t.Clients, c)
	go menu.LoggedInMenu(c)
	go t.SendMessageEveryone(c)
}
