package router

import "github.com/fanjindong/payne/msg"

type IRouter map[msg.Tag]Handler
