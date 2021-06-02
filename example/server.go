package main

import (
	"context"
	"github.com/fanjindong/payne"
	"github.com/fanjindong/payne/msg"
)

func main() {
	s := payne.NewTcpServer(payne.WithRouter(map[msg.Tag]payne.Handler{
		Ping: PingHandler,
	}))
	err := s.Start()
	if err != nil {
		panic(err)
	}
}

const (
	Ping msg.Tag = iota
)

func PingHandler(ctx context.Context, req payne.IRequest) (payne.IReply, error) {
	text := string(req.GetData())
	return payne.NewReply(req.GetConn(), msg.NewMsg(Ping, []byte("是的，"+text))), nil
}
