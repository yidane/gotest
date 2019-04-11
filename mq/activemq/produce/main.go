package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"github.com/yidane/gotest/mq/activemq"
	"strconv"
	"time"
)

func main() {
	conn := activemq.Dail()
	defer activemq.Disconnect(conn)

	const total = 1000000

	now := time.Now()

	for i := 0; i < total; i++ {
		body := []byte(strconv.Itoa(i))
		err := commonSend(conn, body)
		if err != nil {
			err = commonSend(conn, body)
			if err != nil {
				err = commonSend(conn, body)
				if err != nil {
					err = commonSend(conn, body)
					if err != nil {
						fmt.Println(err)
						continue
					}
				}
			}
		}

		fmt.Println(i)

		//err = transactionSend(conn, body)
		//if err != nil {
		//	log.Fatal(err)
		//}
	}

	t := time.Now().Sub(now).Nanoseconds()
	fmt.Println("total", t)
	fmt.Println("avg", float64(t)/float64(total))
}

func commonSend(conn *stomp.Conn, body []byte) error {
	return conn.Send(activemq.Topic, activemq.ContentType, body)
}

func transactionSend(conn *stomp.Conn, body []byte) error {
	tran := conn.Begin()
	err := tran.Send(activemq.Topic, activemq.ContentType, body)
	if err != nil {
		return err
	}

	return tran.Commit()
}
