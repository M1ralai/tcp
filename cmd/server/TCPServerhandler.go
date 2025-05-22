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
	c := client.NewClient(*u, conn)
	t.Clients = append(t.Clients, *c)
	t.serverMessages(c.User.Username + " logged in\n")
	go menu.LoggedInMenu(c)
	go t.SendMessageEveryone(*c)
	go t.clientRequestHandler(*c)
}
