package activemq

import (
	"github.com/go-stomp/stomp"
	"log"
)

const (
	Topic       = "topic"
	ContentType = "yidane"
)

func Dail() *stomp.Conn {
	conn, err := stomp.Dial("tcp", "172.17.0.2:61613")
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func Disconnect(conn *stomp.Conn) {
	err := conn.Disconnect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("active message queue exited")
}
