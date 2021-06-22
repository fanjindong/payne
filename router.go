package payne

import (
	"context"
	"github.com/fanjindong/payne/msg"
)

type IRouter interface {
	Before(context.Context, IRequest) error
	Handler(context.Context, IRequest) error
	After(context.Context, IRequest) error
}

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
