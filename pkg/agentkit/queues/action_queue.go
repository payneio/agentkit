package queues

import (
	"agentkit/pkg/agentkit/datatypes"
	"container/list"
	"math/rand"
	"time"
	// "nanomsg.org/go/mangos/protocol/pub"
	// "nanomsg.org/go/mangos/protocol/sub"
	// "nanomsg.org/go/mangos/v2"
)

// func server(url string) {
// 	var sock mangos.Socket
// 	var err error
// 	if sock, err = pub.NewSocket(); err != nil {
// 		die("can't get new pub socket: %s", err)
// 	}
// 	if err = sock.Listen(url); err != nil {
// 		die("can't listen on pub socket: %s", err.Error())
// 	}
// 	for {
// 		// Could also use sock.RecvMsg to get header
// 		d := date()
// 		fmt.Printf("SERVER: PUBLISHING DATE %s\n", d)
// 		if err = sock.Send([]byte(d)); err != nil {
// 			die("Failed publishing: %s", err.Error())
// 		}
// 		time.Sleep(time.Second)
// 	}
// }

// func client(url string, name string) {
// 	var sock mangos.Socket
// 	var err error
// 	var msg []byte

// 	if sock, err = sub.NewSocket(); err != nil {
// 		die("can't get new sub socket: %s", err.Error())
// 	}
// 	if err = sock.Dial(url); err != nil {
// 		die("can't dial on sub socket: %s", err.Error())
// 	}
// 	// Empty byte array effectively subscribes to everything
// 	err = sock.SetOption(mangos.OptionSubscribe, []byte(""))
// 	if err != nil {
// 		die("cannot subscribe: %s", err.Error())
// 	}
// 	for {
// 		if msg, err = sock.Recv(); err != nil {
// 			die("Cannot recv: %s", err.Error())
// 		}
// 		fmt.Printf("CLIENT(%s): RECEIVED %s\n", name, string(msg))
// 	}
// }

// ActionQueue is a FIFO queue for actions. An action queue linearizes actions
// from a mind and feeds them into the actuator dispatcher.
type ActionQueue interface {
	Peek() *datatypes.Action
	Enqueue(*datatypes.Action)
	Dequeue() *datatypes.Action
	Clear()
}

type InMemoryActionQueue struct {
	q *list.List
}

func NewInMemoryActionQueue() *InMemoryActionQueue {
	return &InMemoryActionQueue{
		q: list.New(),
	}
}

func (q *InMemoryActionQueue) Peek() *datatypes.Action {

	// This is a lowbrow way to add some time to consumer `for` loops.
	// This will be replaced with a proper event bus.
	time.Sleep(time.Duration(900+rand.Intn(200)) * time.Millisecond)

	action := q.q.Front()
	if action == nil {
		return nil
	}

	return action.Value.(*datatypes.Action)
}

func (q *InMemoryActionQueue) Enqueue(p *datatypes.Action) {
	q.q.PushBack(p)
}

func (q *InMemoryActionQueue) Dequeue() *datatypes.Action {
	action := q.q.Front()
	if action == nil {
		return nil
	}
	q.q.Remove(action)
	return action.Value.(*datatypes.Action)
}

func (q *InMemoryActionQueue) Clear() {
	q.q = list.New()
}
