package menu

import (
	"fmt"

	"www.github/M1ralai/tcp/cmd/client"
	"www.github/M1ralai/tcp/cmd/message"
)

func LoggedInMenu(c client.Client) {
	writeLoginData(c)
	for {
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
		}
	}
}

func writeLoginData(c client.Client) {
	c.Conn.Write([]byte("\033[2J\033[1;1H"))
	c.Conn.Write([]byte("You logged in as: "))
	c.Conn.Write([]byte(c.User.Username))
	c.Conn.Write([]byte("    :logout for log out\n"))
}
