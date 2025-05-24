package server

import (
	"fmt"

	"www.github/M1ralai/tcp/cmd/chatroom"
	"www.github/M1ralai/tcp/cmd/client"
	"www.github/M1ralai/tcp/cmd/connio"
	"www.github/M1ralai/tcp/cmd/users"
)

var listClients = "1"
var listOptions = "2"
var createCR = "3"
var listCR = "4"
var enterCR = "5"

func (t *TCPServer) clientRequestHandler(c client.Client) {
	for {
		req := <-c.Req
		switch req {
		case listClients:
			t.listClients(c)
		case listOptions:
			t.listOptions(c)
		case createCR:
			t.createCR(c)
		case listCR:
			t.listCR(c)
		case enterCR:
			t.enterCR(c)
		}
	}
}

func (t *TCPServer) listClients(c client.Client) {
	for i := range users.Uusers {
		if users.Uusers[i].IsLoggedIn {
			c.Conn.Write([]byte(t.Clients[i].User.Username))
			c.Conn.Write([]byte(" is online"))
			c.Conn.Write([]byte("\n"))
		}
	}
}

func (t *TCPServer) listOptions(c client.Client) {
	c.Conn.Write([]byte(":logout for close the connection and log out user \n"))
	c.Conn.Write([]byte(":list for list currently connected users \n"))
	c.Conn.Write([]byte(":createcr for creating chat room \n"))
	c.Conn.Write([]byte(":listcr for list currently connected users \n"))
}

func (t *TCPServer) listCR(c client.Client) {
	for i := range t.chatRooms {
		write := t.chatRooms[i].Name
		c.Conn.Write([]byte(write))
		c.Conn.Write([]byte("\n"))
	}
}

func (t *TCPServer) createCR(c client.Client) {
	c.Conn.Write([]byte("What you wanna give a name of a chatroom: \n"))
	name, err := connio.Read(c.Conn)
	if err != nil {
		fmt.Printf("connection lost from a client %+v \n", c.Conn.RemoteAddr())
		c.Conn.Close()
		return
	}
	c.Conn.Write([]byte("What you wanna give a password to a chatroom: \n"))
	password, err := connio.Read(c.Conn)
	if err != nil {
		fmt.Printf("connection lost from a client %+v \n", c.Conn.RemoteAddr())
		c.Conn.Close()
		return
	}
	for i := range t.chatRooms {
		if t.chatRooms[i].Name == name || name == "default" {
			c.RS <- false
			return
		}
	}
	cr := chatroom.NewChatroom(name, c, password, len(t.chatRooms))
	t.chatRooms = append(t.chatRooms, cr)
	c.RS <- true
}

func (t *TCPServer) enterCR(c client.Client) {
	t.listCR(c)
	c.Conn.Write([]byte("Type the name of you wanted to connect chatroom \n"))
	c.Conn.Write([]byte("(just press enter in password section when you try to connect default server) \n"))
	CRname, err := connio.Read(c.Conn)
	if err != nil {
		fmt.Printf("connection lost from a client %+v \n", c.Conn.RemoteAddr())
		c.Conn.Close()
		return
	}
	for i := range t.chatRooms {
		if CRname == t.chatRooms[i].Name {
			c.Conn.Write([]byte("Type the password of this chatroom \n"))
			CRpassword, err := connio.Read(c.Conn)
			if err != nil {
				fmt.Printf("connection lost from a client %+v \n", c.Conn.RemoteAddr())
				c.Conn.Close()
				return
			}
			if CRpassword == t.chatRooms[i].Password {
				var uarr []client.Client
				for j := range t.chatRooms[c.Room].Users {
					if t.chatRooms[c.Room].Users[j].User.Username != c.User.Username {
						uarr = append(uarr, t.chatRooms[c.Room].Users[j])
					}
				}
				t.chatRooms[c.Room].Users = uarr
				t.chatRooms[i].Users = append(t.chatRooms[i].Users, c)
				c.Room = i
				c.RS <- true
				return
			}
		}
	}
}
