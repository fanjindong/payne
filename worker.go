package payne

import (
	"context"
	"fmt"
	"github.com/fanjindong/payne/msg"
)

type IWorker interface {
	Start()
	Close()
	Do(c Conn, m msg.IMsg) error
}

type Worker struct {
	r     IRouter
	close chan bool
	queue chan IRequest
}

func NewWorker(router IRouter) *Worker {
	return &Worker{r: router, queue: make(chan IRequest, 1024), close: make(chan bool)}
}

func (w Worker) Start() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-w.close:
					return
				case req := <-w.queue:
					w.do(req)
				}
			}
		}()
	}
}

func (w Worker) Do(c Conn, m msg.IMsg) error {
	w.queue <- NewRequest(c, m)
	return nil
}

func (w Worker) do(req IRequest) error {
	return w.r[req.GetTag()](context.Background(), req)
}

func (w Worker) Close() {
	fmt.Println("worker close")
	close(w.close)
	close(w.queue)
}
