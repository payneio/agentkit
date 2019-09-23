package queues

import (
	"agentkit/pkg/agentkit/datatypes"
	"container/list"
)

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
