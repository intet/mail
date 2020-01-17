package main

import (
	"context"
	pb "github.com/intet/mail/sender-service/proto/mail"
	"log"
)

type handler struct {
}

func (h *handler) Process(ctx context.Context, msg *pb.ErrorMsg) error {
	log.Println(msg.Msg);
	return nil
}
