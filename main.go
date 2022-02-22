package main

import (
	"GoCausal/graph"
	"GoCausal/utils"
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

	q := utils.LinkedQueue{}
	q.Append("1")
	q.Append("2")
	q.Append("3")
	fmt.Println(q)
	fmt.Println(q.Contains("4"))
	fmt.Println(q.Pop())
	fmt.Println(q)
}
