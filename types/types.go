package types

type Key string

type VectorClock map[string]uint64

type Value struct {
	Data[] byte
	Clock  VectorClock
}