package node

type Cluster struct {
    Nodes map[string]*Node  
}

func NewCluster(nodes []*Node) *Cluster {
    c := &Cluster{
        Nodes: make(map[string]*Node),
    }

    for _, n := range nodes {
        c.Nodes[n.NodeID] = n
    }

    return c
}

func (c *Cluster) AddNode(node *Node){
    c.Nodes[node.NodeID] = node
}

func (c *Cluster) GetNode(nodeID string) *Node{
    return c.Nodes[nodeID]
}

func (c *Cluster) GetNodes() *[]Node{
    nodeArr := make([], 0, len(c.Nodes))
    for _, n in c.Nodes{
        append(nodeArr, n)
    }
    return nodeArr
}

func (c *Cluster) GetNodeIDs() *[]Node{
    nodeArr := make([], 0, len(c.Nodes))
    for i, _ in c.Nodes{
        append(nodeArr, i)
    }
    return nodeArr
}