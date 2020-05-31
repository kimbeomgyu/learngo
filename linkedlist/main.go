package main

import "fmt"

// Node is node
type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

// AddNode is O(1)
func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
}

// RemoveNode is O(N)
func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}

	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
	}
	node.next = nil
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.printNodes()

	list.RemoveNode(list.root.next)

	list.printNodes()

	list.RemoveNode(list.root)

	list.printNodes()

	list.RemoveNode(list.tail)

	list.printNodes()

	list.RemoveNode(list.root.next.next)

	list.printNodes()
}

func (l *LinkedList) printNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Println(node.val)
}
