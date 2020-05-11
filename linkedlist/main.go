package main

import "fmt"

// Node is node
type Node struct {
	next *Node
	val  int
}

func addNode(root *Node, val int) {
	var tail *Node
	tail = root

	for tail.next != nil {
		tail = tail.next
	}

	node := &Node{val: val}
	tail.next = node

}

func main() {
	var root *Node

	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		addNode(root, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Println(node.val)

}
