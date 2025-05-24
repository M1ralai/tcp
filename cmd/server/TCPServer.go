package server

import (
	"fmt"
	"log"
	"net"

	"www.github/M1ralai/tcp/cmd/chatroom"
	"www.github/M1ralai/tcp/cmd/client"
)

type TCPServer struct {
	TCPaddr       string
	Clients       []client.Client
	chatRooms     []chatroom.Chatroom
	listener      net.Listener
	serverMessage chan string
}

func NewTCPServer(TCPAddr string) *TCPServer {
	return &TCPServer{
		TCPaddr:       TCPAddr,
		chatRooms:     make([]chatroom.Chatroom, 0),
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
	t.chatRooms = append(t.chatRooms, chatroom.DefaultChatroom())
	go t.acceptLoop()
	return nil
}

func (t *TCPServer) SendMessage(c client.Client) {
	for {
		msg := <-c.Msg
<<<<<<< HEAD
		for i := range t.chatRooms[c.Room].Users {
			if t.chatRooms[c.Room].Users[i].Conn.RemoteAddr() == c.Conn.RemoteAddr() {
				t.chatRooms[c.Room].Users[i].Conn.Write([]byte("\033[1A\033[K"))
			}
			t.chatRooms[c.Room].Users[i].Conn.Write(msg.Time)
			t.chatRooms[c.Room].Users[i].Conn.Write([]byte("    "))
			t.chatRooms[c.Room].Users[i].Conn.Write([]byte(msg.SenderName))
			t.chatRooms[c.Room].Users[i].Conn.Write([]byte(" :  sended:   "))
			t.chatRooms[c.Room].Users[i].Conn.Write(msg.Msg)
			t.chatRooms[c.Room].Users[i].Conn.Write([]byte("\n"))
=======
		for i := range t.Clients {
			t.Clients[i].Conn.Write([]byte("\033[1A\033[K"))
			t.Clients[i].Conn.Write(msg.Time)
			t.Clients[i].Conn.Write([]byte("    "))
			t.Clients[i].Conn.Write([]byte(msg.SenderName))
			t.Clients[i].Conn.Write([]byte(" :  sended:   "))
			t.Clients[i].Conn.Write(msg.Msg)
			t.Clients[i].Conn.Write([]byte("\n"))
>>>>>>> master
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

func (t *TCPServer) serverMessages(msg string) {
	for i := range t.Clients {
		t.Clients[i].Conn.Write([]byte(msg))
	}
}
