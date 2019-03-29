package main

import (
	"fmt"
	"github.com/Tencent/bk-cmdb/src/framework/core/log"
	"github.com/streadway/amqp"
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

	ch, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan int)

	count := 0
	go func() {
		for d := range ch {
			fmt.Println(d.MessageCount)
			count++
		}
	}()

	<-forever
	fmt.Println(count)
}
