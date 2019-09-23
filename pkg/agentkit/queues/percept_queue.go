package queues

import (
	"agentkit/pkg/agentkit/datatypes"
	"container/list"
)

// PerceptQueue is a FIFO queue for percepts
type PerceptQueue interface {
	Peek() *datatypes.Percept
	Enqueue(*datatypes.Percept)
	Dequeue() *datatypes.Percept
	Clear()
}

type InMemoryPerceptQueue struct {
	q *list.List
}

func NewInMemoryPerceptQueue() *InMemoryPerceptQueue {
	return &InMemoryPerceptQueue{
		q: list.New(),
	}
}

func (q *InMemoryPerceptQueue) Peek() *datatypes.Percept {
	if q.q.Front() == nil {
		return nil
	}
	return q.q.Front().Value.(*datatypes.Percept)
}

func (q *InMemoryPerceptQueue) Enqueue(p *datatypes.Percept) {
	q.q.PushBack(p)
}

func (q *InMemoryPerceptQueue) Dequeue() *datatypes.Percept {
	p := q.q.Front()
	if q.q.Front() == nil {
		return nil
	}
	q.q.Remove(p)
	return p.Value.(*datatypes.Percept)
}

func (q *InMemoryPerceptQueue) Clear() {
	q.q = list.New()
}
