package queues

import (
	"agentkit/pkg/agentkit/datatypes"
	"container/list"
	"math/rand"
	"time"
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

	// This is a lowbrow way to add some time to consumer `for` loops.
	// This will be replaced with a proper event bus.
	time.Sleep(time.Duration(900+rand.Intn(200)) * time.Millisecond)

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
