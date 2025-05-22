package client

import (
	"net"

	"www.github/M1ralai/tcp/cmd/message"
	"www.github/M1ralai/tcp/cmd/users"
)

type Client struct {
	User users.Users
	Conn net.Conn
	Msg  chan message.Message
	Req  chan (string)
}

func NewClient(u users.Users, conn net.Conn) *Client {
	return &Client{
		User: u,
		Conn: conn,
		Msg:  make(chan message.Message),
		Req:  make(chan string),
	}
}
