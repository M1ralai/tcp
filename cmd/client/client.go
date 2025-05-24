package client

import (
	"net"

	"www.github/M1ralai/tcp/cmd/message"
	"www.github/M1ralai/tcp/cmd/users"
)

type Client struct {
	User users.Users
	Conn net.Conn
	Room int
	Msg  chan message.Message
	Req  chan (string)
	RS   chan (bool)
}

func NewClient(u users.Users, conn net.Conn) *Client {
	return &Client{
		User: u,
		Conn: conn,
		Room: 0,
		Msg:  make(chan message.Message),
		Req:  make(chan string),
		RS:   make(chan bool),
	}
}
