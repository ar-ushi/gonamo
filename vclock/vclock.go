package vclock

//list of tubples [(nid1, tick1), (nid2, tick2), (...)] where pid is the process and tick is the clock value
type VClock map[string]uint64

func newClock() VClock{
	return make(VClock)
}

func (v *VClock) Increment (nodeId string){
	v[nodeId] = v[nodeId] + 1
}

func (v *VClock) Get(nodeId string) uint64{
	return v[nodeId]
}

func (v *VClock) Merge(v2 VClock) VClock{
	res := newClock()
	for nodeId, version := range v{
		res[nodeId] = version
	}
	for nodeId, version := range v2{
		if res[nodeId] < version {
			res[nodeId] = version
		}
	}
	return res
}

/** returns true if v2 is descendant of v1*/ 
func (v1 *VClock) DescendantOf(v2 VClock) bool{
	for nodeId, version := range v2 {
		if v1[nodeId] < version {
			return false
		}
	}
	return true
}

func (v1 *VClock) Equals(v2 VClock) bool{
for nodeId, version := range v2 {
		if v1[nodeId] == version {
			return true
		}
	}
	return false
}

