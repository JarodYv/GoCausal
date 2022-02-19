package graph

type NodeType int32
type NodeVariableType int32

const (
	MEASURED  NodeType = 1
	LATENT    NodeType = 2
	ERROR     NodeType = 3
	SESSION   NodeType = 4
	RANDOMIZE NodeType = 5
	LOCK      NodeType = 6
	NO_TYPE   NodeType = 7
)

const (
	DOMAIN              NodeVariableType = 1
	INTERVENTION_STATUS NodeVariableType = 2
	INTERVENTION_VALUE  NodeVariableType = 3
)

type INode interface {
	GetName() string
	SetName(string)
	GetNodeType() NodeType
	SetNodeType(NodeType)
	GetNodeVariableType() NodeVariableType
	SetNodeVariableType(NodeVariableType)
	GetCenterX() int
	SetCenterX(int)
	GetCenterY() int
	SetCenterY(int)
	SetCenter(int, int)
}

type Node struct {
	INode
	Attribute
	name     string
	nodeType NodeType
	varType  NodeVariableType
	centerX  int
	centerY  int
}

func (node *Node) GetName() string {
	return node.name
}

func (node *Node) SetName(name string) {
	node.name = name
}

func (node *Node) GetNodeType() NodeType {
	return node.nodeType
}

func (node *Node) SetNodeType(nodeType NodeType) {
	node.nodeType = nodeType
}

func (node *Node) GetNodeVariableType() NodeVariableType {
	return node.varType
}

func (node *Node) SetNodeVariableType(varType NodeVariableType) {
	node.varType = varType
}

func (node *Node) GetCenterX() int {
	return node.centerX
}

func (node *Node) SetCenterX(x int) {
	node.centerX = x
}

func (node *Node) GetCenterY() int {
	return node.centerY
}

func (node *Node) SetCenterY(y int) {
	node.centerY = y
}

func (node *Node) SetCenter(x, y int) {
	node.centerX = x
	node.centerY = y
}

func (node *Node) eq(n *Node) bool {
	return node.name == n.name
}
