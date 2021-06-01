package router

import (
	"github.com/fanjindong/payne/conn"
	"github.com/fanjindong/payne/msg"
)

type IReply interface {
	GetConn() conn.IConn
	msg.IMsg
}

type Reply struct {
	conn conn.IConn
	msg.IMsg
}

func NewReply(conn conn.IConn, IMsg msg.IMsg) *Reply {
	return &Reply{conn: conn, IMsg: IMsg}
}

func (r Reply) GetConn() conn.IConn {
	return r.conn
}
