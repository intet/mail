package main;

import (
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/kafka"
	pb "github.com/intet/mail/sender-service/proto/mail"

	"fmt"
)
const topic = "mail.send"
const errorTopic = "error_topic"
const workerCount = 3

func main() {
	srv := micro.NewService(
		micro.Name("sender.service.mail"),
	)
	srv.Init()

	errorPublisher := micro.NewPublisher(errorTopic, srv.Client())

	mailChan := make(chan *pb.Msg)
	defer close(mailChan)
	for w := 1; w <= workerCount; w++ {
		go worker(mailChan, errorPublisher)
	}
	micro.RegisterSubscriber(topic, srv.Server(), &handler{mailChan: mailChan})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}



