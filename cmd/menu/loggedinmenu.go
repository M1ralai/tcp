package menu

import (
	"fmt"

	"www.github/M1ralai/tcp/cmd/client"
	"www.github/M1ralai/tcp/cmd/connio"
	"www.github/M1ralai/tcp/cmd/message"
)

func LoggedInMenu(c *client.Client) {
	writeLoginData(c)
	for {
<<<<<<< HEAD
		fmt.Print(c.Room)
		i, err := connio.Read(c.Conn)
		if err != nil {
			fmt.Printf("connection lost from a client %+v \n", c.Conn.RemoteAddr())
			c.User.LogOut()
			return
		}
		if i[0] == ':' {
			switch i {
			case ":list":
				c.Req <- "1"
			case ":options":
				c.Req <- "2"
			case ":createcr":
				c.Req <- "3"
				//RS stands for request status
				RS := <-c.RS
				if RS {
					c.Conn.Write([]byte("chatroom created! \n"))
				} else {
					c.Conn.Write([]byte("there is already a chatroom with that name \n"))
				}
			case ":listcr":
				c.Req <- "4"
			case ":entercr":
				c.Req <- "5"
				RS := <-c.RS
				if RS {
					c.Conn.Write([]byte("entered that chatroom! \n"))
				} else {
					c.Conn.Write([]byte("chatroom password or name is incorrect \n"))
				}
			case ":logout":
				c.User.LogOut()
				c.Conn.Close()
				return
			}
		} else {
			msg := message.NewMessage([]byte(i), c.User.Username)
			c.Msg <- msg
			continue
=======
		i, err := read(c.Conn)
		if err == nil {
			if i[0] != ':' {
				msg := message.NewMessage([]byte(i), c.User.Username)
				c.Msg <- msg
				continue
			} else {
				switch i {
				case ":logout":
					c.User.LogOut()
					c.Conn.Close()
					return
				}
			}
		} else {
			fmt.Printf("connection lost from a client %+v \n", c.Conn.RemoteAddr())
			c.User.LogOut()
			return
>>>>>>> master
		}
	}
}

func writeLoginData(c *client.Client) {
	c.Conn.Write([]byte("\033[2J\033[1;1H"))
	c.Conn.Write([]byte("type :option for see options"))
	c.Conn.Write([]byte("\n"))
	c.Conn.Write([]byte("You logged in as: "))
	c.Conn.Write([]byte(c.User.Username))
	c.Conn.Write([]byte("    :logout for log out\n"))
	c.Conn.Write([]byte(" h/ m/ s/ \n"))
}
