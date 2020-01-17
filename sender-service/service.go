package main

import (
	"net/smtp"
	"fmt"
	pb "github.com/intet/mail/sender-service/proto/mail"
)

type service struct {
}

const DEFAULT_HOST = "smtp.yandex.ru"
const DEFAULT_PORT = "587";

func (s *service) sendMail(msg *pb.Msg) error {

	auth := smtp.PlainAuth(msg.From, msg.From, msg.Password, DEFAULT_HOST)

	if err := smtp.SendMail(DEFAULT_HOST+":"+DEFAULT_PORT, auth, msg.From, msg.Hdrs, msg.Body); err != nil {
		fmt.Println("error send mail: ", err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}
