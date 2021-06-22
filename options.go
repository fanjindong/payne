package payne

import "github.com/fanjindong/payne/codec"

type IOption func(o *Option)

type Option struct {
	codec  codec.ICodec
	router IRouter
}

func NewOption() *Option {
	return &Option{codec: &codec.LvCodec{}}
}

func WithCodec(c codec.ICodec) IOption {
	return func(o *Option) { o.codec = c }
}

func WithRouter(r IRouter) IOption {
	return func(o *Option) { o.router = r }
}
