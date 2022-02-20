package graph

import "fmt"

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

func (e *Edge) Exchange() {
	node := *e.node2
	*e.node2 = *e.node1
	*e.node1 = node

	endpoint := e.endpoint2
	e.endpoint2 = e.endpoint1
	e.endpoint1 = endpoint
}

func (e *Edge) SetEndPoint1(endpoint Endpoint) {
	e.endpoint1 = endpoint
	if PointingLeft(e.endpoint1, e.endpoint2) {
		e.Exchange()
	}
}

func (e *Edge) SetEndPoint2(endpoint Endpoint) {
	e.endpoint2 = endpoint
	if PointingLeft(e.endpoint1, e.endpoint2) {
		e.Exchange()
	}
}

/*
GetProximalEndpoint

return the endpoint nearest to the given node; returns NULL if the given node is not along the edge.
*/
func (e *Edge) GetProximalEndpoint(node *Node) Endpoint {
	if e.node1.Equals(node) {
		return e.endpoint1
	} else if e.node2.Equals(node) {
		return e.endpoint2
	} else {
		return NULL
	}
}

/*
GetDistalEndpoint

return the endpoint furthest from the given node
*/
func (e *Edge) GetDistalEndpoint(node *Node) Endpoint {
	if e.node1.Equals(node) {
		return e.endpoint2
	} else if e.node2.Equals(node) {
		return e.endpoint1
	} else {
		return NULL
	}
}

/*
GetDistalNode

Given one node along the edge, returns the node at the opposite end of the edge
*/
func (e *Edge) GetDistalNode(node *Node) *Node {
	if e.node1.Equals(node) {
		return e.node2
	} else if e.node2.Equals(node) {
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

func (e *Edge) Equals(edge *Edge) bool {
	return e.endpoint1 == edge.endpoint1 && e.endpoint2 == edge.endpoint2 &&
		e.node1.Equals(edge.node1) && e.node2.Equals(edge.node2)
}

func (e *Edge) LessThan(edge *Edge) bool {
	return e.node1.LessThan(edge.node1) || e.node2.LessThan(edge.node2)
}

func (e *Edge) ToString() string {
	edgeString := e.node1.GetName() + " "
	if e.endpoint1 == TAIL {
		edgeString += "-"
	} else if e.endpoint1 == ARROW {
		edgeString += "<"
	} else {
		edgeString += "o"
	}
	edgeString += "-"
	if e.endpoint2 == TAIL {
		edgeString += "-"
	} else if e.endpoint2 == ARROW {
		edgeString += ">"
	} else {
		edgeString += "o"
	}
	edgeString += " "
	edgeString += e.node2.GetName()
	return edgeString
}

/*
PointingLeft

Check whether the edge is [A <-- B] or [A <-o B] point direction
*/
func PointingLeft(endpoint1, endpoint2 Endpoint) bool {
	return endpoint1 == ARROW && (endpoint2 == TAIL || endpoint2 == CIRCLE)
}

func NewEdge(node1, node2 *Node, end1, end2 Endpoint) (*Edge, error) {
	var edge *Edge
	if node1 == nil || node2 == nil {
		return edge, fmt.Errorf("nodes must not be nil")
	}
	if end1 == NULL || end2 == NULL {
		return edge, fmt.Errorf("endpoints must not be NULL")
	}
	if PointingLeft(end1, end2) {
		edge = &Edge{
			node1:     node2,
			node2:     node1,
			endpoint1: end2,
			endpoint2: end1,
		}
	} else {
		edge = &Edge{
			node1:     node1,
			node2:     node2,
			endpoint1: end1,
			endpoint2: end2,
		}
	}
	return edge, nil
}
