package main

import "fmt"

// Node
// it will contain a val and next node address
// val is int
type Node struct {
	Val  int
	Next *Node
}

// Linked List
// initialy head is empty
type LinkedList struct {
	head *Node
}

// adding an element in linked list
func (list *LinkedList) AddNode(val int) {
	// check head is empty or not,
	// if it is empty then assign it new node, otherwise iterate in list
	// and check for node which has a nil point for next
	if list.head == nil {
		list.head = &Node{
			Val:  val,
			Next: nil,
		}
	} else {
		temp := list.head

		for temp.Next != nil {
			temp = temp.Next
		}

		temp.Next = &Node{
			Val:  val,
			Next: nil,
		}
	}
}

// display all nodes
func (list *LinkedList) DisplayNodes() {
	temp := list.head

	for temp != nil {
		printMe(temp.Val)
		temp = temp.Next
	}
}

func printMe(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	// creating a int slice
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	list := LinkedList{}

	for i := range arr {
		list.AddNode(arr[i])
	}

	list.DisplayNodes()
}
