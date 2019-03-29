package main

import (
	"fmt"
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

		log.Println("active message queue exited")
	}()

	sub, err := conn.Subscribe("testTopic", stomp.AckAuto)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for m := range sub.C {
			fmt.Println(string(m.Body))
		}
	}()

	time.Sleep(time.Minute)
}
