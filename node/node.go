package node
import (
"github.com/ar-ushi/gonamo/storageEngine"
 "github.com/ar-ushi/gonamo/types"
 "github.com/ar-ushi/gonamo/vclock"
)

// what does a node need to have at bare minimum?

// id + address + position in the ring + allows data items to be added

type Node struct {
	NodeID string
	Addr string
	Storage *StorageEngine
}

func NewNode(nodeID string, addr string) *Node{
	return &Node{
		NodeID:      nodeID,
		Addr:        addr,
		Storage:     storageEngine.newStorageEngine(),		
	}
}

func (n *Node) Get(key types.Key) ([]types.Value, bool){
	return n.Storage.Get(key)
}

func (n *Node)  Put(key types.Key, value string, ctx vclock.VClock){
var clock vclock.VClock
	if ctx == nil {
		clock = vclock.NewVClock()
	} else {
		clock = ctx.Copy()
	}
	clock.Increment(n.NodeID)

	value := types.Value{
		Data:  val,
		Clock: clock,
	}

	n.Storage.Put(key, value)
}
