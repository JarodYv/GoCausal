package graph

import "gonum.org/v1/gonum/mat"

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
	ExistsDirectedPathFromTo(Node, Node) bool
	ExistsUndirectedPathFromTo(Node, Node) bool
	ExistsSemidirectedPathFromTo(Node, Node) bool
	ExistsInducingPath(Node, Node) bool
	ExistsTrek(Node, Node) bool
	FullyConnect(Endpoint)
	ReorientAllWith(Endpoint)
	GetAdjacentNodes(Node) []Node
	GetAncestors([]Node) []Node
	GetChildren(Node) []Node
	GetParents(*Node) []*Node
	GetConnectivity()
	GetDescendants([]Node) []Node
	GetEdge(Node, Node) Edge
	GetDirectedEdge(Node, Node) Edge
	GetNodeEdges(Node) []Edge
	GetConnectingEdges(Node, Node) []Edge
	GetGraphEdges() []Edge
	GetEndpoint(Node, Node)
	GetInDegree(Node) int
	GetOutDegree(Node) int
	GetDegree(Node) int
	GetNode(string) Node
	GetNodes() []Node
	GetNodeNames() []string
	GetNumEdges() int
	GetNumConnectedEdges() int
	GetNumNodes() int
	IsAdjacentTo(Node, Node) bool
	IsAncestorOf(Node, Node) bool
	PossibleAncestor(Node, Node) bool
	IsChildOf(Node, Node) bool
	IsParentOf(Node, Node) bool
	IsProperAncestorOf(Node, Node) bool
	IsProperDescendantOf(Node, Node) bool
	IsDescendantOf(Node, Node) bool
	DefNonDescendent(Node, Node)
	IsDefNonCollider(Node, Node, Node) bool
	IsDefCollider(Node, Node, Node) bool
	IsDConnectedTo(Node, Node, Node) bool
	IsDSeparatedFrom(Node, Node, Node) bool
	MaybeDConnectedTo(Node, Node, Node) bool
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

func (g *Graph) Clear() {
	g.nodes = []*Node{}
	g.varNum = 0
	g.nodeMap = map[*Node]int{}
	g.graph.Reset()
	g.dPath.Reset()
}
