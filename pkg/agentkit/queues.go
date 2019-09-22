package agentkit

import "container/list"

// PerceptQueue is a FIFO queue for percepts
type PerceptQueue interface {
	Peek() *Percept
	Enqueue(*Percept)
	Dequeue() *Percept
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

func (q *InMemoryPerceptQueue) Peek() *Percept {
	if q.q.Front() == nil {
		return nil
	}
	return q.q.Front().Value.(*Percept)
}

func (q *InMemoryPerceptQueue) Enqueue(p *Percept) {
	q.q.PushBack(p)
}

func (q *InMemoryPerceptQueue) Dequeue() *Percept {
	p := q.q.Front()
	if q.q.Front() == nil {
		return nil
	}
	q.q.Remove(p)
	return p.Value.(*Percept)
}

func (q *InMemoryPerceptQueue) Clear() {
	q.q = list.New()
}

// ActionQueue is a FIFO queue for actions. An action queue linearizes actions
// from a mind and feeds them into the actuator dispatcher.
type ActionQueue interface {
	Peek() *Action
	Enqueue(*Action)
	Dequeue() *Action
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

func (q *InMemoryActionQueue) Peek() *Action {
	action := q.q.Front()
	if action == nil {
		return nil
	}
	return action.Value.(*Action)
}

func (q *InMemoryActionQueue) Enqueue(p *Action) {
	q.q.PushBack(p)
}

func (q *InMemoryActionQueue) Dequeue() *Action {
	action := q.q.Front()
	if action == nil {
		return nil
	}
	q.q.Remove(action)
	return action.Value.(*Action)
}

func (q *InMemoryActionQueue) Clear() {
	q.q = list.New()
}
