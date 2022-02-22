package graph

import "GoCausal/utils"

type NodePoint struct {
	node *Node
	edge *Edge
}

func MapKeyInNodeSlice(haystack []*Node, needle *Node) bool {
	set := make(map[*Node]struct{})
	for _, e := range haystack {
		set[e] = struct{}{}
	}
	_, ok := set[needle]
	return ok
}

func (n *NodePoint) GetDistalNode() *Node {
	return n.edge.GetDistalNode(n.node)
}

func ExistsDirectedPathFromToBreadthFirst(nodeFrom, nodeTo *Node, g *Graph) bool {
	v := []*Node{nodeFrom}
	q := utils.LinkedQueue{}
	q.Append(nodeFrom)
	for q.Size() > 0 {
		t := q.Pop().(*Node)
		for _, u := range g.GetAdjacentNodes(t) {
			if g.IsParentOf(t, u) && g.IsParentOf(u, t) {
				return true
			}
			edge := g.GetEdge(t, u)
			c := TraverseDirected(t, edge)
			if c == nil || MapKeyInNodeSlice(v, c) {
				continue
			}
			if c == nodeTo {
				return true
			}
			v = append(v, c)
			q.Append(c)
		}
	}
	return false
}

/*
IsDConnectedTo

Returns true if node1 is d-connected to node2 on the set of nodes z.
*/
func IsDConnectedTo(node1, node2 *Node, z []*Node, g *Graph) bool {
	// TODO
	if node1 == node2 {
		return true
	}
	q := utils.LinkedQueue{}
	for _, e := range g.GetNodeEdges(node1) {
		if e.GetDistalNode(node1) == node2 {
			return true
		}
		q.Append(NodePoint{node: node1, edge: e})
	}
	for q.Size() > 0 {
		nodePoint := q.Pop().(NodePoint)
		a := nodePoint.node
		b := nodePoint.GetDistalNode()
		for _, e := range g.GetNodeEdges(b) {
			c := e.GetDistalNode(b)
			if c == a {
				continue
			}
			if Reachable(nodePoint.edge, e, a, z, g) {
				if c == node2 {
					return true
				} else {
					q.Append(NodePoint{node: b, edge: e})
				}
			}
		}
	}
	return false
}

/*
Reachable

Determines if two edges do or do not form a block for d-separation,
conditional on a set of nodes z starting from a node a
*/
func Reachable(edge1, edge2 *Edge, a *Node, z []*Node, g *Graph) bool {
	b := edge1.GetDistalNode(a)
	collider := edge1.GetProximalEndpoint(b) == ARROW && edge2.GetProximalEndpoint(b) == ARROW
	if !collider && !MapKeyInNodeSlice(z, b) {
		return true
	}
	ancestor := IsAncestor(b, z, g)
	return collider && ancestor
}

/*
IsAncestor

Determines if a given node is an ancestor of any node in a set of nodes z.
*/
func IsAncestor(node *Node, z []*Node, g *Graph) bool {
	if MapKeyInNodeSlice(z, node) {
		return true
	}
	q := utils.LinkedQueue{}
	for _, n := range z {
		q.Append(n)
	}
	for q.Size() > 0 {
		t := q.Pop().(*Node)
		if t == node {
			return true
		}
		for _, c := range g.GetParents(t) {
			if !q.Contains(c) {
				q.Append(c)
			}
		}
	}
	return false
}
