package main

import (
	"context"
	pb "github.com/intet/mail/sender-service/proto/mail"
)

type handler struct {
	mailService service
}
func (h *handler) Process(ctx context.Context, msg *pb.Msg) error {
	return h.mailService.sendMail(msg)
}
