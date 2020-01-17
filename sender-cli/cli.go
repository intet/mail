package main

import (
	"log"
	"context"
	pb "github.com/intet/mail/sender-service/proto/mail"

	"github.com/micro/go-micro"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const (
	defaultFilename = "msg.json"
)

const topic = "mail.send"
const errorTopic = "error_topic"

const DEFAULT_FILE = "mail.json"

func main() {

	// Set up a connection to the server.
	srv := micro.NewService(micro.Name("sender.cli"))
	srv.Init()

	publisher := micro.NewPublisher(topic, srv.Client())
	micro.RegisterSubscriber(errorTopic, srv.Server(), &handler{})

	sendMail(publisher)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func sendMail(publisher micro.Publisher) {
	msg := &pb.Msg{From: "someone@mail.com", Password: "***", Hdrs: []string{"anotherOne"}, Body: []byte("Hello world")}
	err := publisher.Publish(context.TODO(), msg)
	if err != nil {
		log.Fatalf("Could not publish mail: %v", err)
	}
	msg, err = parseFile(DEFAULT_FILE)
	if err != nil {
		log.Fatalf("Could not parse files: %v", err)
	}
	err = publisher.Publish(context.TODO(), msg)
	if err != nil {
		log.Fatalf("Could not publish mail: %v", err)
	}
}

func parseFile(file string) (*pb.Msg, error) {
	var consignment *pb.Msg
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}