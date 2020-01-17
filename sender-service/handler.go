package main

import (
	"context"
	"github.com/micro/go-micro"
	pb "github.com/intet/mail/sender-service/proto/mail"
)

type handler struct {
	mailService    *service
	errorPublisher micro.Publisher
}
func (h *handler) Process(ctx context.Context, msg *pb.Msg) error {
	err := h.mailService.sendMail(msg)

	if err != nil {
		errMsg := &pb.ErrorMsg{Msg: err.Error()}
		h.errorPublisher.Publish(context.TODO(), errMsg)
	}

	return err
}
