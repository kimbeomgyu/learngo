package main

import "fmt"

// Node is node
type Node struct {
	next *Node
	val  int
}

// AddNode is O(1)
func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

// RemoveNode is O(N)
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	printNodes(root)

	root, tail = RemoveNode(root.next, root, tail)

	printNodes(root)

	root, tail = RemoveNode(root, root, tail)

	printNodes(root)

	root, tail = RemoveNode(tail, root, tail)

	printNodes(root)

	root, tail = RemoveNode(tail, root, tail)

	printNodes(root)
}

func printNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Println(node.val)
	fmt.Println("root: ", root.val)
	fmt.Println("tail: ", node.val)
}
