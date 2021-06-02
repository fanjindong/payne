package payne

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

type IServer interface {
	Start() error
	Stop() error
	SetRouter(r IRouter)
}

type TcpServer struct {
	port     int
	listener net.Listener
	option   *Option
}

func NewTcpServer(ops ...IOption) *TcpServer {
	o := NewOption()
	for _, op := range ops {
		op(o)
	}
	s := &TcpServer{option: o}
	return s
}

func (s *TcpServer) Start() error {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	s.listener = l
	defer s.Stop()

	fmt.Println("server start")
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, os.Kill)

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				continue
			}
			go NewConn(c, s.option.router).Start()
		}
	}()

	<-exit
	return nil
}

func (s TcpServer) Stop() error {
	fmt.Println("server stop")
	return s.listener.Close()
}
