package datatypes

import "time"

// Percept is a datatype required for the percept queue
type Percept struct {
	Label string
	Data  interface{}
	TS    time.Time
}

// Action is an action datatype
type Action struct {
	Label string
	Data  interface{}
	TS    time.Time
}

type AgentCoordinator struct {
	Name        string
	LastCheckin time.Time
	Status      string
}

type Agent struct {
	Name        string
	Address     string
	Coordinator AgentCoordinator
}
