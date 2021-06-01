package router

import (
	"github.com/fanjindong/payne/conn"
	"github.com/fanjindong/payne/msg"
)

type IRequest interface {
	GetConn() conn.IConn
	msg.IMsg
}

type Request struct {
	conn conn.IConn
	msg.IMsg
}

func NewRequest(conn conn.IConn, IMsg msg.IMsg) *Request {
	return &Request{conn: conn, IMsg: IMsg}
}

func (r Request) GetConn() conn.IConn {
	return r.conn
}
