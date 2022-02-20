package graph

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Triple struct {
	Node1 *Node
	Node2 *Node
	Node3 *Node
}

type IGraph interface {
	AddBidirectedEdge(*Node, *Node)
	AddDirectedEdge(*Node, *Node)
	AddUndirectedEdge(*Node, *Node)
	AddNondirectedEdge(*Node, *Node)
	AddPartiallyOrientedEdge(*Node, *Node)
	AddEdge(*Edge) bool
	AddNode(*Node) bool
	Clear()
	ContainsEdge(*Edge) bool
	ContainsNode(*Node) bool
	ExistsDirectedCycle() bool
	ExistsDirectedPathFromTo(*Node, *Node) bool
	ExistsUndirectedPathFromTo(*Node, *Node) bool
	ExistsSemidirectedPathFromTo(*Node, *Node) bool
	ExistsInducingPath(*Node, *Node) bool
	ExistsTrek(*Node, *Node) bool
	FullyConnect(Endpoint)
	ReorientAllWith(Endpoint)
	GetAdjacentNodes(*Node) []*Node
	GetAncestors([]*Node) []*Node
	GetChildren(*Node) []*Node
	GetParents(*Node) []*Node
	GetConnectivity()
	GetDescendants([]*Node) []*Node
	GetEdge(*Node, *Node) *Edge
	GetDirectedEdge(*Node, *Node) *Edge
	GetNodeEdges(*Node) []*Edge
	GetConnectingEdges(*Node, *Node) []*Edge
	GetGraphEdges() []*Edge
	GetEndpoint(*Node, *Node)
	GetInDegree(*Node) int
	GetOutDegree(*Node) int
	GetDegree(*Node) int
	GetNode(string) *Node
	GetNodes() []*Node
	GetNodeNames() []string
	GetNumEdges() int
	GetNumConnectedEdges(*Node) int
	GetNumNodes() int
	IsAdjacentTo(*Node, *Node) bool
	IsAncestorOf(*Node, *Node) bool
	PossibleAncestor(*Node, *Node) bool
	IsChildOf(*Node, *Node) bool
	IsParentOf(*Node, *Node) bool
	IsProperAncestorOf(*Node, *Node) bool
	IsProperDescendantOf(*Node, *Node) bool
	IsDescendantOf(*Node, *Node) bool
	DefNonDescendent(*Node, *Node)
	IsDefNonCollider(*Node, *Node, *Node) bool
	IsDefCollider(*Node, *Node, *Node) bool
	IsDConnectedTo(*Node, *Node, *Node) bool
	IsDSeparatedFrom(*Node, *Node, *Node) bool
	MaybeDConnectedTo(*Node, *Node, *Node) bool
	IsPattern() bool
	SetPattern(bool)
	IsPAG() bool
	SetPAG(bool)
	IsDirectedFromTo(Node, Node) bool
	IsUndirectedFromTo(Node, Node) bool
	DefVisible(Edge) bool
	IsExogenous(Node) bool
	GetNodesInto(Node, Endpoint) []Node
	GetNodesOutOf(Node, Endpoint) []Node
	RemoveEdge(Edge)
	RemoveConnectingEdge(Node, Node)
	RemoveConnectingEdges(Node, Node)
	RemoveEdges([]Edge)
	RemoveNode(Node)
	RemoveNodes([]Node)
	SetEndpoint(Node, Node, Endpoint)
	Subgraph([]Node) IGraph
	TransferNodesAndEdges(IGraph)
	TransferAttributes(IGraph)
	GetAmbiguousTriples()
	GetUnderlines()
	GetDottedUnderlines()
	IsAmbiguousTriple(Node, Node, Node) bool
	IsUnderlineTriple(Node, Node, Node) bool
	IsDottedUnderlineTriple(Node, Node, Node) bool
	AddAmbiguousTriple(Node, Node, Node)
	AddUnderlineTriple(Node, Node, Node)
	AddDottedUnderlineTriple(Node, Node, Node)
	RemoveAmbiguousTriple(Node, Node, Node)
	RemoveUnderlineTriple(Node, Node, Node)
	RemoveDottedUnderlineTriple(Node, Node, Node)
	SetAmbiguousTriple(Node, Node, Node)
	SetUnderlineTriple(Node, Node, Node)
	SetDottedUnderlineTriple(Node, Node, Node)
	GetCausalOrdering()
	IsParameterizable(Node) bool
	IsTimeLagModel() bool
	GetSepset(Node, Node) []Node
	SetNodes([]Node)
}

type Graph struct {
	IGraph
	Attribute
	nodes                  []*Node
	nodeMap                map[*Node]int
	varNum                 int
	graph                  mat.Dense
	dPath                  mat.Dense
	ambiguousTriples       []Triple
	underlineTriples       []Triple
	dottedUnderlineTriples []Triple
	pattern                bool
	pag                    bool
}

func (g *Graph) adjustDPath(i, j int) {
	g.dPath.Set(i, j, 1)
	for k := 0; k < len(g.nodes); k++ {
		if g.dPath.At(i, k) == 1 {
			g.dPath.Set(j, k, 1)
		}
		if g.dPath.At(k, j) == 1 {
			g.dPath.Set(k, i, 1)
		}
	}
}

func (g *Graph) reconstituteDPath(edges []*Edge) {
	for i := 0; i < len(g.nodes); i++ {
		g.adjustDPath(i, i)
	}
	for _, edge := range edges {
		node1 := edge.GetNode1()
		node2 := edge.GetNode2()
		i := g.nodeMap[node1]
		j := g.nodeMap[node2]
		g.adjustDPath(i, j)
	}
}

func MapKeyInNodeSlice(haystack []*Node, needle *Node) bool {
	set := make(map[*Node]struct{})
	for _, e := range haystack {
		set[e] = struct{}{}
	}
	_, ok := set[needle]
	return ok
}

func (g *Graph) collectAncestors(node *Node, ancestors []*Node) {
	if MapKeyInNodeSlice(ancestors, node) {
		return
	}
	ancestors = append(ancestors, node)
	parents := g.GetParents(node)
	if parents != nil {
		for _, p := range parents {
			g.collectAncestors(p, ancestors)
		}
	}
}

/*
AddDirectedEdge

Adds a directed edge --> to the graph.
*/
func (g *Graph) AddDirectedEdge(node1, node2 *Node) {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	g.graph.Set(j, i, 1)
	g.graph.Set(i, j, -1)

	g.adjustDPath(i, j)
}

/*
AddEdge

Adds the specified edge to the graph, provided it is not already in the# graph.
*/
func (g *Graph) AddEdge(edge *Edge) bool {
	node1 := edge.GetNode1()
	node2 := edge.GetNode2()
	endpoint1 := edge.GetEndpoint1()
	endpoint2 := edge.GetEndpoint2()

	i := g.nodeMap[node1]
	j := g.nodeMap[node2]

	e1 := g.graph.At(i, j)
	e2 := g.graph.At(j, i)

	bidirected := e2 == 1 && e1 == 1
	existingEdge := !bidirected && (e2 != 0 || e1 != 0)

	if endpoint1 == TAIL {
		if existingEdge {
			return false
		}
		if endpoint2 == TAIL {
			if bidirected {
				g.graph.Set(j, i, float64(TAIL_AND_ARROW))
				g.graph.Set(i, j, float64(TAIL_AND_ARROW))
			} else {
				g.graph.Set(j, i, float64(TAIL))
				g.graph.Set(i, j, float64(TAIL))
			}
		} else if endpoint2 == ARROW {
			if bidirected {
				g.graph.Set(j, i, float64(ARROW_AND_ARROW))
				g.graph.Set(i, j, float64(TAIL_AND_ARROW))
			} else {
				g.graph.Set(j, i, float64(ARROW))
				g.graph.Set(i, j, float64(TAIL))
			}
			g.adjustDPath(i, j)
		} else if endpoint2 == CIRCLE {
			if bidirected {
				return false
			} else {
				g.graph.Set(j, i, float64(CIRCLE))
				g.graph.Set(i, j, float64(TAIL))
			}
		} else {
			return false
		}
	} else if endpoint1 == ARROW {
		if endpoint2 == ARROW {
			if existingEdge {
				if e1 == 2 || e2 == 2 {
					return false
				}
				if g.graph.At(j, i) == float64(ARROW) {
					g.graph.Set(j, i, float64(ARROW_AND_ARROW))
				} else {
					g.graph.Set(j, i, float64(TAIL_AND_ARROW))
				}
				if g.graph.At(i, j) == float64(ARROW) {
					g.graph.Set(i, j, float64(ARROW_AND_ARROW))
				} else {
					g.graph.Set(i, j, float64(TAIL_AND_ARROW))
				}
			} else {
				g.graph.Set(j, i, float64(ARROW))
				g.graph.Set(i, j, float64(ARROW))
			}
		} else {
			return false
		}
	} else if endpoint1 == CIRCLE {
		if existingEdge {
			return false
		}
		if endpoint2 == ARROW {
			if bidirected {
				return false
			} else {
				g.graph.Set(j, i, float64(ARROW))
				g.graph.Set(i, j, float64(CIRCLE))
			}

		} else if endpoint2 == CIRCLE {
			if bidirected {
				return false
			} else {
				g.graph.Set(j, i, float64(CIRCLE))
				g.graph.Set(i, j, float64(CIRCLE))
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}

/*
AddNode

Adds a node to the graph.
Precondition: The proposed name of the node cannot already be used by any other node in the same graph.
*/
func (g *Graph) AddNode(node *Node) bool {
	if MapKeyInNodeSlice(g.nodes, node) {
		return false
	}
	g.nodes = append(g.nodes, node)
	g.nodeMap[node] = g.varNum
	g.varNum += 1

	return true
}

/*
Clear

Removes all nodes (and therefore all edges) from the graph.
*/
func (g *Graph) Clear() {
	g.nodes = []*Node{}
	g.varNum = 0
	g.nodeMap = map[*Node]int{}
	g.graph.Reset()
	g.dPath.Reset()
}

/*
ContainsEdge

Determines whether this graph contains the given edge.
Returns true iff the graph contain 'edge'.
*/
func (g *Graph) ContainsEdge(edge *Edge) bool {
	endpoint1 := edge.GetEndpoint1()
	endpoint2 := edge.GetEndpoint2()

	node1 := edge.GetNode1()
	node2 := edge.GetNode2()

	i := g.nodeMap[node1]
	j := g.nodeMap[node2]

	e1 := Endpoint(g.graph.At(i, j))
	e2 := Endpoint(g.graph.At(j, i))

	if endpoint1 == TAIL {
		if endpoint2 == TAIL {
			return (e2 == TAIL && e1 == TAIL) || (e2 == TAIL_AND_ARROW && e1 == TAIL_AND_ARROW)
		} else if endpoint2 == ARROW {
			return (e1 == TAIL && e2 == ARROW) || (e1 == TAIL_AND_ARROW && e2 == ARROW_AND_ARROW)
		} else if endpoint2 == CIRCLE {
			return e1 == TAIL && e2 == CIRCLE
		} else {
			return false
		}
	} else if endpoint1 == ARROW {
		if endpoint2 == ARROW {
			return (e1 == ARROW && e2 == ARROW) || (e1 == TAIL_AND_ARROW && e2 == TAIL_AND_ARROW) || (e1 == ARROW_AND_ARROW || e2 == ARROW_AND_ARROW)
		} else {
			return false
		}
	} else if endpoint1 == CIRCLE {
		if endpoint2 == ARROW {
			return e1 == CIRCLE && e2 == ARROW
		} else if endpoint2 == CIRCLE {
			return e1 == CIRCLE && e2 == CIRCLE
		} else {
			return false
		}
	} else {
		return false
	}
}

/*
ContainsNode

Determines whether this graph contains the given node.
Returns true iff the graph contains 'node'.
*/
func (g *Graph) ContainsNode(node *Node) bool {
	if node == nil {
		return false
	}
	_, ok := g.nodeMap[node]
	return ok
}

/*
ExistsDirectedCycle

Returns true iff there is a directed cycle in the graph.
*/
func (g *Graph) ExistsDirectedCycle() bool {
	for _, node := range g.nodes {
		if ExistsDirectedPathFromToBreadthFirst(node, node, g) {
			return true
		}
	}
	return false
}

/*
ExistsTrek

Returns true iff a trek exists between two nodes in the graph.
A trek exists if there is a directed path between the two nodes or else,
for some third node in the graph, there is a path to each of the two nodes in question.
*/
func (g *Graph) ExistsTrek(node1, node2 *Node) bool {
	for _, node := range g.nodes {
		if g.IsAncestorOf(node, node1) && g.IsAncestorOf(node, node2) {
			return true
		}
	}
	return false
}

/*
Equals

Determines whether this graph is equal to some other graph,
in the sense that they contain the same nodes and the sets of edges defined over these
nodes in the two graphs are isomorphic typewise.
That is, if node A and B exist in both graphs, and if there are, e.g., three edges between A and B
in the first graph, two of which are directed edges and one of which is an undirected edge,
then in the second graph there must also be two directed edges and one undirected edge between nodes A and B.
*/
func (g *Graph) Equals(graph *Graph) bool {
	// TODO
	return false
}

/*
GetAdjacentNodes

Returns a slice of nodes adjacent to the given node.
*/
func (g *Graph) GetAdjacentNodes(node *Node) []*Node {
	j := g.nodeMap[node]
	var adjNodes []*Node
	for i := 0; i < g.varNum; i++ {
		if g.graph.At(j, i) != 0 && g.graph.At(i, j) != 0 {
			n := g.nodes[i]
			adjNodes = append(adjNodes, n)
		}
	}
	return adjNodes
}

/*
GetParents

Return the list of parents of a node.
*/
func (g *Graph) GetParents(node *Node) []*Node {
	j := g.nodeMap[node]
	var parents []*Node
	for i := 0; i < g.varNum; i++ {
		e1 := Endpoint(g.graph.At(i, j))
		e2 := Endpoint(g.graph.At(j, i))
		if (e1 == TAIL && e2 == ARROW) || (e1 == TAIL_AND_ARROW && e2 == ARROW_AND_ARROW) {
			n := g.nodes[i]
			parents = append(parents, n)
		}
	}
	return parents
}

/*
GetAncestors

Returns a slice of ancestors for the given nodes.
*/
func (g *Graph) GetAncestors(nodes []*Node) []*Node {
	var ancestors []*Node
	for _, n := range nodes {
		g.collectAncestors(n, ancestors)
	}
	return ancestors
}

/*
GetChildren

Returns a slice of children for a node.
*/
func (g *Graph) GetChildren(node *Node) []*Node {
	i := g.nodeMap[node]
	var children []*Node
	for j := 0; j < g.varNum; j++ {
		e1 := Endpoint(g.graph.At(i, j))
		e2 := Endpoint(g.graph.At(j, i))
		if (e1 == TAIL && e2 == ARROW) || (e1 == TAIL_AND_ARROW && e2 == ARROW_AND_ARROW) {
			n := g.nodes[i]
			children = append(children, n)
		}
	}
	return children
}

/*
GetInDegree

Returns the number of arrow endpoints adjacent to the node.
*/
func (g *Graph) GetInDegree(node *Node) int {
	i := g.nodeMap[node]
	inDegree := 0
	for j := 0; j < g.varNum; j++ {
		e := Endpoint(g.graph.At(i, j))
		if e == ARROW {
			inDegree += 1
		} else if e == ARROW_AND_ARROW {
			inDegree += 2
		}
	}
	return inDegree
}

/*
GetOutDegree

Returns the number of null endpoints adjacent to the node.
*/
func (g *Graph) GetOutDegree(node *Node) int {
	i := g.nodeMap[node]
	outDegree := 0
	for j := 0; j < g.varNum; j++ {
		e := Endpoint(g.graph.At(i, j))
		if e == TAIL || e == TAIL_AND_ARROW {
			outDegree += 1
		}
	}
	return outDegree
}

/*
GetDegree

Returns the total number of edges into and out of the node.
*/
func (g *Graph) GetDegree(node *Node) int {
	i := g.nodeMap[node]
	degree := 0
	for j := 0; j < g.varNum; j++ {
		e := Endpoint(g.graph.At(i, j))
		if e == ARROW || e == TAIL || e == CIRCLE {
			degree += 1
		} else if e != NULL {
			degree += 2
		}
	}
	return degree
}

/*
GetMaxDegree

Returns the degree of the node with the max degree
*/
func (g *Graph) GetMaxDegree() int {
	max := -1
	for _, node := range g.nodes {
		deg := g.GetDegree(node)
		if deg > max {
			max = deg
		}
	}
	return max
}

/*
GetNode

Returns the node with the given string name.
In case of accidental duplicates, the first node encountered with the given name is returned.
In case no node exists with the given name, nil is returned.
*/
func (g *Graph) GetNode(name string) *Node {
	for _, node := range g.nodes {
		if node.GetName() == name {
			return node
		}
	}
	return nil
}

/*
GetNodes

Returns the list of nodes for the graph.
*/
func (g *Graph) GetNodes() []*Node {
	return g.nodes
}

/*
GetNodeNames

Returns the names of the nodes, in the order of get_nodes.
*/
func (g *Graph) GetNodeNames() []string {
	var names []string
	for _, node := range g.nodes {
		names = append(names, node.GetName())
	}
	return names
}

/*
GetNumEdges

Returns the number of edges in the entire graph.
*/
func (g *Graph) GetNumEdges() int {
	edges := 0
	for i := 0; i < g.varNum; i++ {
		for j := i + 1; j < g.varNum; j++ {
			e := Endpoint(g.graph.At(i, j))
			if e == ARROW || e == TAIL || e == CIRCLE {
				edges += 1
			} else if e != NULL {
				edges += 2
			}
		}
	}
	return edges
}

/*
GetNumConnectedEdges

Returns the number of edges in the graph which are connected to a particular node.
*/
func (g *Graph) GetNumConnectedEdges(node *Node) int {
	edges := 0
	i := g.nodeMap[node]
	for j := 0; j < g.varNum; j++ {
		e := Endpoint(g.graph.At(j, i))
		if e == ARROW || e == TAIL || e == CIRCLE {
			edges += 1
		} else if e != NULL {
			edges += 2
		}
	}
	return edges
}

/*
GetNumNodes

Return the number of nodes in the graph.
*/
func (g *Graph) GetNumNodes() int {
	return g.varNum
}

/*
IsAdjacentTo

Return true iff node1 is adjacent to node2 in the graph.
*/
func (g *Graph) IsAdjacentTo(node1, node2 *Node) bool {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	e := Endpoint(g.graph.At(j, i))
	return e != NULL
}

/*
IsAncestorOf

Return true iff node1 is an ancestor of node2.
*/
func (g *Graph) IsAncestorOf(node1, node2 *Node) bool {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	e := Endpoint(g.graph.At(j, i))
	return e == ARROW
}

/*
IsDescendantOf

Returns true iff node1 is a descendant of node2.
*/
func (g *Graph) IsDescendantOf(node1, node2 *Node) bool {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	e := Endpoint(g.graph.At(i, j))
	return e == ARROW
}

/*
IsChildOf

Return true iff node1 is a child of node2.
*/
func (g *Graph) IsChildOf(node1, node2 *Node) bool {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	e := Endpoint(g.graph.At(i, j))
	return e == ARROW || e == ARROW_AND_ARROW
}

/*
IsParentOf

Returns true iff node1 is a parent of node2.
*/
func (g *Graph) IsParentOf(node1, node2 *Node) bool {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	e := Endpoint(g.graph.At(j, i))
	return e == ARROW || e == ARROW_AND_ARROW
}

/*
IsProperAncestorOf

Returns true iff node1 is a proper ancestor of node2.
*/
func (g *Graph) IsProperAncestorOf(node1, node2 *Node) bool {
	return (g.IsAncestorOf(node1, node2)) && (!node1.Equals(node2))
}

/*
IsProperDescendantOf

Returns true iff node1 is a proper descendant of node2.
*/
func (g *Graph) IsProperDescendantOf(node1, node2 *Node) bool {
	return (g.IsDescendantOf(node1, node2)) && (!node1.Equals(node2))
}

/*
GetEdge

Returns the edge connecting node1 and node2, provided a unique such edge exists.
*/
func (g *Graph) GetEdge(node1, node2 *Node) *Edge {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	e1 := Endpoint(g.graph.At(i, j))
	e2 := Endpoint(g.graph.At(j, i))
	if e1 == NULL {
		return nil
	}
	edge, err := NewEdge(node1, node2, e1, e2)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return edge
}

/*
GetDirectedEdge

Returns the directed edge from node1 to node2, if there is one.
*/
func (g *Graph) GetDirectedEdge(node1, node2 *Node) *Edge {
	i := g.nodeMap[node1]
	j := g.nodeMap[node2]
	e1 := Endpoint(g.graph.At(i, j))
	e2 := Endpoint(g.graph.At(j, i))
	if e1 > ARROW || e1 == NULL || (e1 == TAIL && e2 == TAIL) {
		return nil
	}
	edge, err := NewEdge(node1, node2, e1, e2)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return edge
}

/*
GetNodeEdges

Returns the slice of edges connected to a particular node.
No particular ordering of the edges in the list is guaranteed.
*/
func (g *Graph) GetNodeEdges(node *Node) []*Edge {
	i := g.nodeMap[node]
	var edges []*Edge
	for j := 0; j < g.varNum; j++ {
		n := g.nodes[j]
		e2 := Endpoint(g.graph.At(j, i))
		if e2 == ARROW || e2 == TAIL || e2 == CIRCLE {
			edge := g.GetEdge(node, n)
			edges = append(edges, edge)
		} else {
			var edge1, edge2 *Edge
			edge2, _ = NewEdge(node, n, ARROW, ARROW)
			e1 := Endpoint(g.graph.At(i, j))
			if e2 == TAIL_AND_ARROW && e1 == ARROW_AND_ARROW {
				edge1, _ = NewEdge(node, n, ARROW, TAIL)
			} else if e2 == ARROW_AND_ARROW && e1 == TAIL_AND_ARROW {
				edge1, _ = NewEdge(node, n, TAIL, ARROW)
			} else if e2 == TAIL_AND_ARROW && e1 == TAIL_AND_ARROW {
				edge1, _ = NewEdge(node, n, TAIL, TAIL)
			}
			if edge1 != nil && edge2 != nil {
				edges = append(edges, edge1)
				edges = append(edges, edge2)
			}
		}
	}
	return edges
}

func (g *Graph) GetGraphEdges() []*Edge {
	var edges []*Edge
	for i := 0; i < g.varNum; i++ {
		node1 := g.nodes[i]
		for j := i + 1; j < g.varNum; j++ {
			node2 := g.nodes[j]
			e2 := Endpoint(g.graph.At(j, i))
			if e2 == ARROW || e2 == TAIL || e2 == CIRCLE {
				edge := g.GetEdge(node1, node2)
				edges = append(edges, edge)
			} else {
				var edge1, edge2 *Edge
				edge2, _ = NewEdge(node1, node2, ARROW, ARROW)
				e1 := Endpoint(g.graph.At(i, j))
				if e2 == TAIL_AND_ARROW && e1 == ARROW_AND_ARROW {
					edge1, _ = NewEdge(node1, node2, ARROW, TAIL)
				} else if e2 == ARROW_AND_ARROW && e1 == TAIL_AND_ARROW {
					edge1, _ = NewEdge(node1, node2, TAIL, ARROW)
				} else if e2 == TAIL_AND_ARROW && e1 == TAIL_AND_ARROW {
					edge1, _ = NewEdge(node1, node2, TAIL, TAIL)
				}
				if edge1 != nil && edge2 != nil {
					edges = append(edges, edge1)
					edges = append(edges, edge2)
				}
			}
		}
	}
	return edges
}
