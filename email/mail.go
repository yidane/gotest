package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"time"
)

var (
	host = "smtp.aliyun.com"
	port = 25

	count int64 = 1
)

type Content struct {
	From    string
	To      []string
	Subject string
	Body    string
}

type MailServer struct {
	c   chan Content
	s   chan os.Signal
	t   chan int
	f   int32
	d   *time.Ticker
	lis *net.Listener
}

func New(l int) *MailServer {
	if l < 0 {
		l = 0
	}

	server := MailServer{
		c: make(chan Content, l),
		s: make(chan os.Signal, 0),
		t: make(chan int, 3), //控制发送邮件线程的数量为3
		f: 0,
		d: time.NewTicker(time.Second),
	}

	server.t <- 1
	server.t <- 2
	server.t <- 3

	return &server
}

func (server *MailServer) Listen(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server.lis = &lis

	http.HandleFunc("/sendEmail", func(writer http.ResponseWriter, request *http.Request) {
		server.addEmail(writer, request)
	})

	atomic.SwapInt32(&server.f, 1)
	log.Println("开启邮件服务")

	return nil
}

func (server *MailServer) Server() {
	go func() {
		err := http.Serve(*server.lis, nil)
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			select {
			case email := <-server.c:
				go server.sendEmail(email, <-server.t)
			case <-server.d.C:
			}
		}
	}()

	server.waitClose()
}

func (server *MailServer) waitClose() {
	signal.Notify(server.s, os.Interrupt, os.Kill)
	s := <-server.s
	close(server.c)
	atomic.SwapInt32(&server.f, 0)
	log.Println(s)
}

func (server *MailServer) sendEmail(content Content, id int) {
	dialer := gomail.NewDialer(host, port, "certificater@aliyun.com", "")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "certificater@aliyun.com")
	msg.SetHeader("To", "15201261252@139.com")
	msg.SetHeader("Subject", "报警邮件"+content.Subject)
	msg.SetBody("text/plain", "报警提醒信息")

	err := dialer.DialAndSend(msg)
	if err != nil {
		fmt.Println(id, content.Subject, err)
	} else {
		fmt.Println(id, content.Subject, "succeed")
	}

	server.t <- id //执行完毕，释放协程
}

func (server *MailServer) running() bool {
	return atomic.LoadInt32(&server.f) == 1
}

func (server *MailServer) addEmail(res http.ResponseWriter, req *http.Request) {
	if !server.running() {
		res.WriteHeader(200)
		_, err := res.Write([]byte("Service is closed"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	if req.Method != "POST" {
		res.WriteHeader(200)
		_, err := res.Write([]byte("未知请求方式"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	c := atomic.LoadInt64(&count)
	atomic.AddInt64(&count, 1)
	server.c <- Content{Subject: strconv.FormatInt(c, 10)}

	res.WriteHeader(200)
	_, err := res.Write([]byte("OK"))
	if err != nil {
		log.Println(err)
	}
}
