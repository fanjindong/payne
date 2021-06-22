package main

import (
	"context"
	"github.com/fanjindong/payne"
	"github.com/fanjindong/payne/codec"
	"github.com/fanjindong/payne/msg"
)

func main() {
	s := payne.NewTcpServer(payne.WithRouter(&PingRouter{}), payne.WithCodec(&codec.LvCodec{}))
	err := s.Start()
	if err != nil {
		panic(err)
	}
}

type PingRouter struct {
}

func (p PingRouter) Before(ctx context.Context, request payne.IRequest) error {
	return nil
}

func (p PingRouter) Handler(ctx context.Context, request payne.IRequest) error {
	text := string(request.GetData())
	request.GetConn().Send(msg.NewMsg([]byte("Pong: " + text)))
	return nil
}

func (p PingRouter) After(ctx context.Context, request payne.IRequest) error {
	return nil
}
