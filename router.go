package payne

import (
	"context"
	"github.com/fanjindong/payne/msg"
)

type Handler func(context.Context, IRequest) (IReply, error)

type IRouter map[msg.Tag]Handler

type IRequest interface {
	GetConn() IConn
	msg.IMsg
}

type Request struct {
	conn IConn
	msg.IMsg
}

func NewRequest(conn IConn, IMsg msg.IMsg) *Request {
	return &Request{conn: conn, IMsg: IMsg}
}

func (r Request) GetConn() IConn {
	return r.conn
}

type IReply interface {
	GetConn() IConn
	msg.IMsg
}

type Reply struct {
	conn IConn
	msg.IMsg
}

func NewReply(conn IConn, IMsg msg.IMsg) *Reply {
	return &Reply{conn: conn, IMsg: IMsg}
}

func (r Reply) GetConn() IConn {
	return r.conn
}
