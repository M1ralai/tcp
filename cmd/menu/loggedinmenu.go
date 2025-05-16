package menu

import (
	"www.github/M1ralai/tcp/cmd/client"
	"www.github/M1ralai/tcp/cmd/message"
)

func LoggedInMenu(c client.Client) {
	writeLoginData(c)
	for {
		i, err := read(c.Conn)
		if err == nil {
			msg := message.NewMessage([]byte(i), c.User.Username)
			c.Msg <- msg
			continue
		} else {
			c.Conn.Close()
		}
	}
}

func writeLoginData(c client.Client) {
	c.Conn.Write([]byte("\033[2J\033[1;1H"))
	c.Conn.Write([]byte("You logged in as: "))
	c.Conn.Write([]byte(c.User.Username))
	c.Conn.Write([]byte("    :logout for log out\n"))
}
