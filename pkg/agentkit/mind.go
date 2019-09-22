package agentkit

import "fmt"

type Mind interface {
	Start()
}

type LoopbackMind struct {
	Percepts PerceptQueue
	Actions  ActionQueue
}

func (m *LoopbackMind) Start() {

	fmt.Println(`Loopback mind is waking.`)

	// agent cycle
	go func(m *LoopbackMind) {
		for {
			if m.Percepts.Peek() != nil {
				percept := m.Percepts.Dequeue()
				action := &Action{
					Label: `echo.` + percept.Label,
					Data:  percept.Data,
				}
				m.Actions.Enqueue(action)
			}
		}
	}(m)

}
