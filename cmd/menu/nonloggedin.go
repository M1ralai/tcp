package menu

import (
	"net"

	"www.github/M1ralai/tcp/cmd/connio"
	"www.github/M1ralai/tcp/cmd/users"
)

func NoLoginMenu(conn net.Conn) *users.Users {
	conn.Write([]byte(":register for create new user\n"))
	conn.Write([]byte(":login for login already exist account\n"))
	var u *users.Users
	var err error
	var input string
	for {
		input, err = connio.Read(conn)
		if err != nil {
			return nil
		}
		switch input {
		case ":register":
			u, err = registerUserMenu(conn)
			if err != nil {
				conn.Write([]byte(err.Error()))
			} else {
				return u
			}
		case ":login":
			u, err = loginUserMenu(conn)
			if err != nil {
				conn.Write([]byte(err.Error()))
			} else {
				return u
			}
		}
	}
}

func registerUserMenu(conn net.Conn) (*users.Users, error) {
	conn.Write([]byte("Give me a username: \n"))
	username, err := connio.Read(conn)
	if err != nil {
		return nil, err
	}
	conn.Write([]byte("Give me a password: \n"))
	password, err := connio.Read(conn)
	if err != nil {
		return nil, err
	}
	u, err := users.RegisterUser(username, password)
	return u, err
}

func loginUserMenu(conn net.Conn) (*users.Users, error) {
	conn.Write([]byte("Give me a username: \n"))
	username, err := connio.Read(conn)
	if err != nil {
		return nil, err
	}
	conn.Write([]byte("Give me a password: \n"))
	password, err := connio.Read(conn)
	if err != nil {
		return nil, err
	}
	u, err := users.LoginUser(username, password)
	//some experiment for git
	//maceraa dolu ameeerika
	return u, err
}
