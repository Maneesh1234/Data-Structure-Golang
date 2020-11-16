package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func (n *Node) AddNode(data int) {
	newNode := Node{data, nil}
	iter := n
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &newNode
}

func (n *Node) displayLinkedlist() {
	temp := n
	for temp != nil {
		fmt.Print(temp.Data, " ")
		temp = temp.Next
	}
	fmt.Println()

}

func (n *Node) deleteFirst() {
	if n != nil && n.Next == nil {
		n = nil
	} else if n != nil {
		*n = *n.Next
	}

}

func (n *Node) deleteEnd() {
	if n != nil && n.Next == nil {
		n = nil
	} else if n != nil {
		iter := n
		for iter.Next.Next != nil {
			iter = iter.Next
		}
		iter.Next = nil
	}

}

func main() {
	newNode := Node{10, nil}
	newNode.AddNode(20)
	newNode.AddNode(30)
	newNode.AddNode(40)
	newNode.AddNode(50)
	newNode.displayLinkedlist()
	newNode.deleteEnd()

	newNode.displayLinkedlist()
}
