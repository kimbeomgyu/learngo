package main

import "fmt"

// Node is node
type Node struct {
	next *Node
	val  int
}

func addNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = addNode(tail, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Println(node.val)

}
