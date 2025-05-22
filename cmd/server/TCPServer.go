package server

import (
	"fmt"
	"log"
	"net"

	"www.github/M1ralai/tcp/cmd/client"
	"www.github/M1ralai/tcp/cmd/users"
)

type TCPServer struct {
	TCPaddr       string
	Clients       []client.Client
	listener      net.Listener
	serverMessage chan string
}

func NewTCPServer(TCPAddr string) *TCPServer {
	return &TCPServer{
		TCPaddr:       TCPAddr,
		serverMessage: make(chan string),
	}
}

func (t *TCPServer) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.TCPaddr)
	if err != nil {
		fmt.Print("error occured while creating listen port: ", err)
		return err
	}
	go t.acceptLoop()
	return nil
}

func (t *TCPServer) SendMessageEveryone(c client.Client) {
	for {
		msg := <-c.Msg
		for i := range t.Clients {
			if t.Clients[i].Conn.RemoteAddr() == c.Conn.RemoteAddr() {
				t.Clients[i].Conn.Write([]byte("\033[1A\033[K"))
			}
			t.Clients[i].Conn.Write(msg.Time)
			t.Clients[i].Conn.Write([]byte("    "))
			t.Clients[i].Conn.Write([]byte(msg.SenderName))
			t.Clients[i].Conn.Write([]byte(" :  sended:   "))
			t.Clients[i].Conn.Write(msg.Msg)
			t.Clients[i].Conn.Write([]byte("\n"))
		}
	}
}

func (t *TCPServer) serverMessages(msg string) {
	for i := range t.Clients {
		t.Clients[i].Conn.Write([]byte(msg))
	}
}

func (t *TCPServer) clientRequestHandler(c client.Client) {
	listClients := "1"
	for {
		req := <-c.Req
		switch req {
		case listClients:
			t.listRequest(c)
		}
	}
}

func (t *TCPServer) listRequest(c client.Client) {
	for i := range users.Uusers {
		if users.Uusers[i].IsLoggedIn {
			c.Conn.Write([]byte(t.Clients[i].User.Username))
			c.Conn.Write([]byte(" is online"))
			c.Conn.Write([]byte("\n"))
		}
	}
}

func (t *TCPServer) acceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Print("error occured while connecting: ", err)
			conn.Close()
			return
		}
		go t.handleCon(conn)
	}
}
