package payne

import (
	"context"
	"github.com/fanjindong/payne/conn"
	"github.com/fanjindong/payne/router"
	"net"
)

type IServer interface {
	Start() error
	Stop() error
	SetRouter(r router.IRouter)
}

type TcpServer struct {
	port   int
	l      net.Listener
	router router.IRouter
}

func NewTcpServer() *TcpServer {
	return &TcpServer{}
}

func (t *TcpServer) Start() error {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	t.l = l
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		pConn := conn.NewConn(c)
		go func() {
			for {
				m, err := pConn.Receive()
				if err != nil {
					break
				}
				reply, err := t.router[m.GetTag()](context.Background(), router.NewRequest(pConn, m))
				if err != nil {
					panic(err)
				}
				if err = pConn.Send(reply); err != nil {
					panic(err)
				}
			}
		}()
	}
	return nil
}

func (t *TcpServer) SetRouter(r router.IRouter) {
	t.router = r
}

func (t TcpServer) Stop() error {
	return t.l.Close()
}
