package graph

type EdgeProperty int32
type Endpoint int32

const (
	dd EdgeProperty = 1
	nl EdgeProperty = 2
	pd EdgeProperty = 3
	pl EdgeProperty = 4
)

const (
	TAIL            Endpoint = -1
	NULL            Endpoint = 0
	ARROW           Endpoint = 1
	CIRCLE          Endpoint = 2
	STAR            Endpoint = 3
	TAIL_AND_ARROW  Endpoint = 4
	ARROW_AND_ARROW Endpoint = 5
)

type Edge struct {
	node1     *Node
	node2     *Node
	endpoint1 Endpoint
	endpoint2 Endpoint
}

func (e *Edge) GetNode1() *Node {
	return e.node1
}

func (e *Edge) GetNode2() *Node {
	return e.node2
}

func (e *Edge) GetEndpoint1() Endpoint {
	return e.endpoint1
}

func (e *Edge) GetEndpoint2() Endpoint {
	return e.endpoint2
}

/*
GetProximalEndpoint

return the endpoint nearest to the given node; returns NULL if the given node is not along the edge.
*/
func (e *Edge) GetProximalEndpoint(node *Node) Endpoint {
	if e.node1.eq(node) {
		return e.endpoint1
	} else if e.node2.eq(node) {
		return e.endpoint2
	} else {
		return NULL
	}
}

func (e *Edge) GetDistalEndpoint(node *Node) Endpoint {
	if e.node1.eq(node) {
		return e.endpoint2
	} else if e.node2.eq(node) {
		return e.endpoint1
	} else {
		return NULL
	}
}

func (e *Edge) GetDistalNode(node *Node) *Node {
	if e.node1.eq(node) {
		return e.node2
	} else if e.node2.eq(node) {
		return e.node1
	} else {
		return nil
	}
}

func (e *Edge) PointsToward(node *Node) bool {
	proximal := e.GetProximalEndpoint(node)
	distal := e.GetDistalEndpoint(node)
	return proximal == ARROW && (distal == TAIL || distal == CIRCLE)
}

/*
PointingLeft

Check whether the edge is [A <-- B] or [A <-o B] point direction
*/
func PointingLeft(endpoint1, endpoint2 Endpoint) bool {
	return endpoint1 == ARROW && (endpoint2 == TAIL || endpoint2 == CIRCLE)
}

func NewEdge(node1, node2 *Node, end1, end2 Endpoint) Edge {
	var edge Edge
	if PointingLeft(end1, end2) {
		edge = Edge{
			node1:     node2,
			node2:     node1,
			endpoint1: end2,
			endpoint2: end1,
		}
	} else {
		edge = Edge{
			node1:     node1,
			node2:     node2,
			endpoint1: end1,
			endpoint2: end2,
		}
	}
	return edge
}
