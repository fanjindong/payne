package conn

import (
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
	c     net.Conn
	codec codec.ICodec
	req   chan msg.IMsg
	reply chan msg.IMsg
}

func NewConn(c net.Conn) *Conn {
	return &Conn{c: c, codec: &codec.TlvCodec{}}
}

func (c Conn) Start() {
	go func() {
		for {
			m, err := c.Receive()
			if err != nil {
				break
			}
			c.req <- m
		}
	}()
	go func() {
		for {
			m := <-c.reply
			if err := c.Send(m); err != nil {
				break
			}
		}
	}()
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
