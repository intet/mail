package main;

import (
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/kafka"
	"fmt"
)
const topic = "mail.send"

func main()  {
	srv := micro.NewService(
		micro.Name("sender.service.mail"),
	)
	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), &handler{})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

