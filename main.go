package main

import (
	"GoCausal/graph"
	"fmt"
)

func changeType(node *graph.Node) {
	node.SetNodeType(graph.LATENT)
}

func main() {
	node1 := graph.Node{}
	node2 := graph.Node{}
	node1.SetName("node1")
	node2.SetName("node2")
	fmt.Println(node1)
	fmt.Println(node2)
	changeType(&node1)
	fmt.Println(node1)
}
