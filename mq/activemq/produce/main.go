package main

import (
	"github.com/go-stomp/stomp"
	"log"
	"time"
)

func main() {
	conn, err := stomp.Dial("tcp", "172.17.0.3:61613")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.Disconnect()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for i := 0; i < 100; i++ {
		err = conn.Send("testTopic", "yidane", []byte(time.Now().String()))
		if err != nil {
			log.Fatal(err)
		}
	}
}
