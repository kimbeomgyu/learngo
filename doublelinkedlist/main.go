package main

import "fmt"

// Node is node
type Node struct {
	next *Node
	prev *Node
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
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

// RemoveNode is O(1)
func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
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
	list.printReverseNodes()

	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
}

func (l *LinkedList) printNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Println(node.val)
}

func (l *LinkedList) printReverseNodes() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Println(node.val)
}
