package ring

import ("github.com/ar-ushi/gonamo/node")


type Ring struct {
	Q int // # of partitions
	Partitions Partition[] // stores hash token, pId, position, primary, replicas
	Version uint64
}


//owns the hash space

func newRing(q int ) *Ring{
	r :=&Ring{Q:q, Partitions: make([]Partition, q)},
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
		r.Partitions[i].Primary = nodeIDs[primaryIndex]
	}
}

func hashKey(key string) uint64 {
	h := sha1.Sum([]byte(key))
	return binary.BigEndian.Uint64(h[:8])
}