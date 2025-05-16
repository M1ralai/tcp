package menu

import (
	"errors"
	"net"

	"www.github/M1ralai/tcp/cmd/users"
)

func NoLoginMenu(conn net.Conn) *users.Users {
	conn.Write([]byte(":register for create new user\n"))
	conn.Write([]byte(":login for login already exist account\n"))
	var u *users.Users
	var err error
	var input string
	for {
		input, err = read(conn)
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
	username, err := read(conn)
	if err != nil {
		return nil, err
	}
	conn.Write([]byte("Give me a password: \n"))
	password, err := read(conn)
	if err != nil {
		return nil, err
	}
	u, err := users.RegisterUser(username, password)
	return u, err
}

func loginUserMenu(conn net.Conn) (*users.Users, error) {
	conn.Write([]byte("Give me a username: \n"))
	username, err := read(conn)
	if err != nil {
		return nil, err
	}
	conn.Write([]byte("Give me a password: \n"))
	password, err := read(conn)
	if err != nil {
		return nil, err
	}
	u, err := users.LoginUser(username, password)
	return u, err
}

func read(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return " ", errors.New("err")
	}
	return string(buf[:(n - 2)]), nil
}
