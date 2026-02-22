package node
import ("github.com/ar-ushi/gonamo/types")

// what does a node need to have at bare minimum?

// id + address + allows data items to be added

type Node struct {
	NodeID string
	Addr string
	Data Item[]
}