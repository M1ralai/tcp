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
	t.chatRooms[0].Enter(*c, "")
	go menu.LoggedInMenu(c)
	go t.SendMessage(*c)
	go t.clientRequestHandler(*c)
}
