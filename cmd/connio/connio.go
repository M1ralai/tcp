package connio

import (
	"net"
)

func Read(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {

		conn.Close()
		return " ", err
	}
	if n == 2 {
		n = 3
	}
	return string(buf[:(n - 2)]), nil
}
