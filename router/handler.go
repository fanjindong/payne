package router

import "context"

type Handler func(context.Context, IRequest) (IReply, error)
