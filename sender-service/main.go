package main;

import (
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/kafka"
	"fmt"
)
const topic = "mail.send"
const errorTopic = "error_topic"

func main()  {
	srv := micro.NewService(
		micro.Name("sender.service.mail"),
	)
	srv.Init()

	errorPublisher := micro.NewPublisher(errorTopic, srv.Client())

	micro.RegisterSubscriber(topic, srv.Server(), &handler{mailService: &service{}, errorPublisher: errorPublisher})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}



