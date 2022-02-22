package graph

/*
BidirectedEdge

constructs a new bidirected edge from node_a to node_b <->
*/
func BidirectedEdge(node1, node2 *Node) *Edge {
	edge := Edge{
		node1:     node1,
		node2:     node2,
		endpoint1: ARROW,
		endpoint2: ARROW,
	}
	return &edge
}

/*
DirectedEdge

constructs a new directed edge from node_a to node_b -->
*/
func DirectedEdge(node1, node2 *Node) *Edge {
	edge := Edge{
		node1:     node1,
		node2:     node2,
		endpoint1: TAIL,
		endpoint2: ARROW,
	}
	return &edge
}

/*
PartiallyOrientedEdge

constructs a new partially oriented edge from node_a to node_b o->
*/
func PartiallyOrientedEdge(node1, node2 *Node) *Edge {
	edge := Edge{
		node1:     node1,
		node2:     node2,
		endpoint1: CIRCLE,
		endpoint2: ARROW,
	}
	return &edge
}

/*
UndirectedEdge

constructs a new undirected edge from node_a to node_b --
*/
func UndirectedEdge(node1, node2 *Node) *Edge {
	edge := Edge{
		node1:     node1,
		node2:     node2,
		endpoint1: TAIL,
		endpoint2: TAIL,
	}
	return &edge
}

/*
IsBidirectedEdge

return true iff an edge is a bidrected edge <->
*/
func IsBidirectedEdge(edge *Edge) bool {
	return edge.GetEndpoint1() == ARROW && edge.GetEndpoint2() == ARROW
}

/*
IsDirectedEdge

return true iff the given edge is a directed edge -->
*/
func IsDirectedEdge(edge *Edge) bool {
	if edge.GetEndpoint1() == TAIL {
		return edge.GetEndpoint2() == ARROW
	} else if edge.GetEndpoint2() == TAIL {
		return edge.GetEndpoint1() == ARROW
	} else {
		return false
	}
}

/*
IsPartiallyOrientedEdge

return true iff the given edge is a partially oriented edge o->
*/
func IsPartiallyOrientedEdge(edge *Edge) bool {
	if edge.GetEndpoint1() == CIRCLE {
		return edge.GetEndpoint2() == ARROW
	} else if edge.GetEndpoint2() == CIRCLE {
		return edge.GetEndpoint1() == ARROW
	} else {
		return false
	}
}

/*
IsUndirectedEdge

return true iff some edge is an undirected edge --
*/
func IsUndirectedEdge(edge *Edge) bool {
	return edge.GetEndpoint1() == TAIL && edge.GetEndpoint2() == TAIL
}

func TraverseDirected(node *Node, edge *Edge) *Node {
	if node == edge.GetNode1() {
		if edge.GetEndpoint1() == TAIL && edge.GetEndpoint2() == ARROW {
			return edge.GetNode2()
		}
	} else if node == edge.GetNode2() {
		if edge.GetEndpoint2() == TAIL && edge.GetEndpoint1() == ARROW {
			return edge.GetNode1()
		}
	}
	return nil
}
