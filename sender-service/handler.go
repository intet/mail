package main

import (
	"context"
	pb "github.com/intet/mail/sender-service/proto/mail"
)

type handler struct {
	mailChan chan<- *pb.Msg
}
func (h *handler) Process(ctx context.Context, msg *pb.Msg) error {
	h.mailChan <- msg
	return nil
}
