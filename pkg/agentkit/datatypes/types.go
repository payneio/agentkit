package datatypes

import "time"

// Percept is a datatype required for the percept queue
type Percept struct {
	Label string
	Data  string
	TS    time.Time
}

// Action is an action datatype
type Action struct {
	Label string
	Data  string
	TS    time.Time
}
