package payne

import (
	"github.com/fanjindong/payne/codec"
	"github.com/fanjindong/payne/msg"
	"net"
)

type IConn interface {
	Start()
	Send(msg.IMsg)
	Receive() msg.IMsg
	Close()
}

type Conn struct {
	c      net.Conn
	codec  codec.ICodec
	worker IWorker
	close  chan bool

	sendQ chan msg.IMsg
}

func NewConn(c net.Conn, codec codec.ICodec, w IWorker) *Conn {
	return &Conn{c: c, codec: codec, worker: w, close: make(chan bool), sendQ: make(chan msg.IMsg, 1)}
}

func (c Conn) Start() {
	go c.send()
	go c.receive()
}

func (c Conn) Send(m msg.IMsg) {
	c.sendQ <- m
}

func (c Conn) send() {
	for {
		select {
		case <-c.close:
			break
		case m := <-c.sendQ:
			if m == nil {
				continue
			}
			data, err := c.codec.Encode(m)
			if err != nil {
				continue
			}
			if _, err = c.c.Write(data); err != nil {
				c.Close()
				return
			}
		}
	}
}

func (c Conn) Receive() msg.IMsg {
	return nil
}

func (c Conn) receive() {
	for {
		select {
		case <-c.close:
			return
		default:
			data, err := c.codec.Decode(c.c)
			if err != nil {
				c.Close()
				return
			}
			c.worker.Do(c, data)
		}
	}
}

func (c Conn) Close() {
	close(c.close)
	close(c.sendQ)
	c.c.Close()
}
