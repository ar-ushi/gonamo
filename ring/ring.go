package ring

import (
    "crypto/sha1"
    "encoding/binary"
    "math"
    "sort"
)

type Ring struct {
	Q int // # of partitions
	Partitions []Partition // stores hash token, pId, position, primary, replicas
	Version uint64
}


type Partition struct {
	ID    int // 0..Q-1
	Token uint64 //start of hash range
	PNode string //nodeId of primary node that owns this
}


//owns the hash space

func newRing(q int ) *Ring{
r := &Ring{
    Q: q,
    Partitions: make([]Partition, q),
}
	maxHash := uint64(math.MaxUint64)
	step := maxHash / uint64(q)

	for i := 0; i < q; i++ {
		r.Partitions[i] = Partition{
			ID:    i,
			Token: uint64(i) * step,
		}
	}

	return r
}

func (r *Ring) AssignNodes(nodeIDs []string){
	sort.Strings(nodeIDs)

	for i := range r.Partitions {
		primaryIndex := i % len(nodeIDs)
		r.Partitions[i].PNode = nodeIDs[primaryIndex]
	}
}

func hashKey(key string) uint64 {
	h := sha1.Sum([]byte(key))
	return binary.BigEndian.Uint64(h[:8])
}

func (r *Ring) lookupPartition(key string) Partition {
	hashedKey := hashKey(key)
	for i := range r.Partitions{
		if (r.Partitions[i].Token >= hashedKey) {
			return r.Partitions[i]
		}
	}
	return r.Partitions[0]
}


func (r *Ring) getPrimaryNode(p Partition) string{
	return p.PNode
}
