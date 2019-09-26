package datatypes

import "time"

// Percept is a datatype required for the percept queue
type Percept struct {
	Label string      `json:"label"`
	Data  interface{} `json:"data"`
	TS    time.Time   `json:"ts"`
}

// Action is an action datatype
type Action struct {
	Label string      `json:"label"`
	Data  interface{} `json:"data"`
	TS    time.Time   `json:"ts"`
}

type Central struct {
	Name        string    `json:"central"`
	Address     string    `json:"address"`
	LastCheckin time.Time `json:"lastCheckin"`
	Status      string    `json:"status"`
}

type Agent struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Central Central `json:"central"`
}

type Belief struct {
	ID        string      `json:"id"`
	Data      interface{} `json:"data"`
	UpdatedAt time.Time   `json:"updatedAt"`
	ChangedAt time.Time   `json:"changedAt"`
}
