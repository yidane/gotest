package main

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://172.17.0.3:5672/")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := channel.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	queue, err := channel.QueueDeclare("testQueue", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	fmt.Println(time.Now())

	count := 0

	go func() {
		for {
			err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
				Type: "text/plain",
				Body: []byte(fmt.Sprintf("hello world eachen,%v", time.Now().UnixNano())),
			})
			if err != nil {
				log.Fatal(err)
			}
			count++
		}
	}()

	<-ctx.Done()
	fmt.Println(time.Now())
	fmt.Println(count)
}
