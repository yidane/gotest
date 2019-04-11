package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"github.com/yidane/gotest/mq/activemq"
	"log"
	"sync/atomic"
	"time"
)

func main() {
	conn := activemq.Dail()
	defer activemq.Disconnect(conn)

	sub, err := conn.Subscribe(activemq.Topic, stomp.AckClient)
	if err != nil {
		log.Fatal(err)
	}

	var total int64

	now := time.Now()
f:
	for {
		select {
		case m := <-sub.C:
			if m.Err != nil {
				err := m.Conn.Nack(m)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			fmt.Println(string(m.Body))
			if m.ShouldAck() {
				err := m.Conn.Ack(m)
				if err != nil {
					log.Println(err)
				}
			}

			atomic.AddInt64(&total, 1)

		case <-time.Tick(time.Second * 10):
			log.Println(sub.Active())
			log.Println("exited")
			break f
		}
	}

	t := time.Now().Sub(now).Nanoseconds()
	fmt.Println("total", t)
	fmt.Println("avg", float64(t)/(1000*float64(total)), "ms")
}

func commonSubscribe(conn *stomp.Conn) (*stomp.Subscription, error) {
	return conn.Subscribe(activemq.Topic, stomp.AckAuto)
}

func clientSubScribe(conn *stomp.Conn) (*stomp.Subscription, error) {
	return conn.Subscribe(activemq.Topic, stomp.AckClient)
}
