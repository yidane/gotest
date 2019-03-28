package main

import (
	"fmt"
	"os"
)

//func sendEmail() {
//	dialer := gomail.NewDialer(host, port, "certificater@aliyun.com", "Aliyun1024.")
//	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
//
//	msg := gomail.NewMessage()
//	msg.SetHeader("From", "certificater@aliyun.com")
//	msg.SetHeader("To", "15201261252@139.com")
//	msg.SetHeader("Subject", "报警邮件")
//	msg.SetBody("text/plain", "报警提醒信息")
//
//	err := dialer.DialAndSend(msg)
//	if err != nil {
//		fmt.Print(err)
//	} else {
//		fmt.Print("succeed")
//	}
//}

const address = ":8123"

func main() {
	mailServer := New(100)
	err := mailServer.Listen(address)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	mailServer.Server()
}
