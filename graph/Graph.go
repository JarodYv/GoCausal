package graph

type IGraph interface {
	AddBidirectedEdge(Node, Node)
	AddDirectedEdge(Node, Node)
	AddUndirectedEdge(Node, Node)
	AddNondirectedEdge(Node, Node)
	AddPartiallyOrientedEdge(Node, Node)
	AddEdge(Edge)
	AddNode(Node)
	Clear()
	ContainsEdge(Edge) bool
	ContainsNode(Node) bool
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
	GetParents(Node) []Node
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
}
