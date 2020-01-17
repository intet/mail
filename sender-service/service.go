package main

import (
	"net/smtp"
	"fmt"
	pb "github.com/intet/mail/sender-service/proto/mail"
	"github.com/micro/go-micro"
	"context"
)

const DEFAULT_HOST = "smtp.yandex.ru"
const DEFAULT_PORT = "587";

func worker(mailChan <-chan *pb.Msg, errorPublisher micro.Publisher) {
	for msg := range mailChan {
		err := sendMail(msg);
		if err != nil {
			errorPublisher.Publish(context.TODO(), &pb.ErrorMsg{Msg: err.Error()})
		}
	}
}

func sendMail(msg *pb.Msg) error {
	auth := smtp.PlainAuth(msg.From, msg.From, msg.Password, DEFAULT_HOST)

	if err := smtp.SendMail(DEFAULT_HOST+":"+DEFAULT_PORT, auth, msg.From, msg.Hdrs, msg.Body); err != nil {
		fmt.Println("error send mail: ", err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}
