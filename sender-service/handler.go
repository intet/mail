package main

import (
	"context"
	pb "github.com/intet/mail/sender-service/proto/mail"
)

type handler struct {

}
func (h *handler) Process(ctx context.Context, msg *pb.Msg) error {

	return nil
}
