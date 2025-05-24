package chatroom

import (
	"net"

	"www.github/M1ralai/tcp/cmd/client"
)

type Chatroom struct {
	ID       int
	Name     string
	Users    []client.Client
	Password string
	admin    net.Conn
}

func NewChatroom(name string, admin client.Client, Password string, ID int) Chatroom {
	var Users []client.Client
	Users = append(Users, admin)
	return Chatroom{
		Name:     name,
		Password: Password,
		admin:    admin.Conn,
		Users:    Users,
		ID:       ID,
	}
}
func DefaultChatroom() Chatroom {
	return Chatroom{
		Name:     "default",
		Password: "",
		admin:    nil,
		Users:    nil,
		ID:       0,
	}
}

func (cr *Chatroom) Enter(c client.Client, p string) bool {
	if cr.Password == p {
		cr.Users = append(cr.Users, c)
		return true
	}
	return false
}
