package menu

import (
	"net"
)

func read(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return " ", err
	}
	return string(buf[:(n - 2)]), nil
}
