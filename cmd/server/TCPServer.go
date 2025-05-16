package server

import (
	"log"
	"net"

	"www.github/M1ralai/tcp/cmd/client"
)

type TCPServer struct {
	TCPaddr  string
	Clients  []client.Client
	listener net.Listener
}

func NewTCPServer(TCPAddr string) *TCPServer {
	return &TCPServer{
		TCPaddr: TCPAddr,
	}
}

func (t *TCPServer) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.TCPaddr)
	if err != nil {
		log.Fatal("error occured while creating listen port: ", err)
		return err
	}
	go t.acceptLoop()
	return nil
}

func (t *TCPServer) SendMessageEveryone(c client.Client) {
	for {
		msg := <-c.Msg
		for i := 0; i < len(t.Clients); i++ {
			if t.Clients[i].User.Username != c.User.Username {
				t.Clients[i].Conn.Write(msg.Time)
				t.Clients[i].Conn.Write([]byte("    "))
				t.Clients[i].Conn.Write([]byte(msg.SenderName))
				t.Clients[i].Conn.Write([]byte(" :  sended:   "))
				t.Clients[i].Conn.Write(msg.Msg)
				t.Clients[i].Conn.Write([]byte("\n"))
			}
		}
	}
}

func (t *TCPServer) acceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Print("error occured while connecting: ", err)
			return
		}
		go t.handleCon(conn)

	}
}
