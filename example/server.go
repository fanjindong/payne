package main

import (
	"context"
	"github.com/fanjindong/payne"
	"github.com/fanjindong/payne/msg"
	"github.com/fanjindong/payne/router"
)

func main() {
	s := payne.NewTcpServer()
	s.SetRouter(map[msg.Tag]router.Handler{
		Ping: PingHandler,
	})
	err := s.Start()
	if err != nil {
		panic(err)
	}
}

const (
	Ping msg.Tag = iota
)

func PingHandler(ctx context.Context, req router.IRequest) (router.IReply, error) {
	text := string(req.GetData())
	return router.NewReply(req.GetConn(), msg.NewMsg(Ping, []byte("是的，"+text))), nil
}
