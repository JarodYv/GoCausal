package graph

import "GoCausal/utils"

func ExistsDirectedPathFromToBreadthFirst(node_from, node_to *Node, g *Graph) bool {
	v := []*Node{node_from}
	q := utils.LinkedQueue{}
	q.Append(node_from)
	// TODO
	return false
}
