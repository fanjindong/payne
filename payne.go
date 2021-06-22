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
	w        IWorker
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
	var err error
	s.listener, err = net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	fmt.Println("server start")
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, os.Kill)

	s.w = NewWorker(s.option.router)
	s.w.Start()

	go func() {
		for {
			c, err := s.listener.Accept()
			if err != nil {
				continue
			}
			NewConn(c, s.option.codec, s.w).Start()
		}
	}()
	defer s.Stop()
	<-exit
	return nil
}

func (s TcpServer) Stop() error {
	fmt.Println("server stop")
	s.listener.Close()
	fmt.Println(s.w)
	s.w.Close()
	return nil
}
