package payne

import (
	"context"
	"github.com/fanjindong/payne/codec"
	"github.com/fanjindong/payne/msg"
	"net"
)

type IConn interface {
	Start()
	Send(msg.IMsg) error
	Receive() (msg.IMsg, error)
	Close()
}

type Conn struct {
	c      net.Conn
	codec  codec.ICodec
	router IRouter
}

func NewConn(c net.Conn, r IRouter) *Conn {
	return &Conn{c: c, codec: &codec.TlvCodec{}, router: r}
}

func (c Conn) Start() {
	for {
		m, err := c.Receive()
		if err != nil {
			break
		}
		reply, err := c.router[m.GetTag()](context.Background(), NewRequest(c, m))
		if err != nil {
			panic(err)
		}
		if err = c.Send(reply); err != nil {
			panic(err)
		}
	}
	return
}

func (c Conn) Send(m msg.IMsg) error {
	data, err := c.codec.Encode(m)
	if err != nil {
		return err
	}
	_, err = c.c.Write(data)
	return err
}

func (c Conn) send(m msg.IMsg) error {
	data, err := c.codec.Encode(m)
	if err != nil {
		return err
	}
	_, err = c.c.Write(data)
	return err
}

func (c Conn) Receive() (msg.IMsg, error) {
	data, err := c.codec.Decode(c.c)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c Conn) Close() {
	c.c.Close()
}
