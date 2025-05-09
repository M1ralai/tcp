package message

import (
	"strconv"
	"time"
)

type Message struct {
	Msg        []byte
	Time       []byte
	SenderName string
}

func NewMessage(msg []byte, SenderName string) Message {
	h, m, s := time.Now().Clock()
	time := []byte(strconv.Itoa(h) + "/" + strconv.Itoa(m) + "/" + strconv.Itoa(s))
	return Message{
		Msg:        msg,
		Time:       time,
		SenderName: SenderName,
	}
}
