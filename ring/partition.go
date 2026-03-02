package ring


type Partition struct {
	ID    int // 0..Q-1
	Token uint64 //start of hash range
	PNode string //nodeId of primary node that owns this
}