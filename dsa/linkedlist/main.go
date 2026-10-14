package main

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

// Beginning mein insert karna - O(1)
func (list *LinkedList[T]) InsertAtBeginning(data T) {
	curr := &Node[T]{
		data: data,
		next: list.head,
	}
	list.head = curr
}

// End mein insert karna - O(N)
func (list *LinkedList[T]) InsertAtEnd(data T) {
	// FIX: Check if head is nil (empty list)
	if list.head == nil {
		list.InsertAtBeginning(data)
		return
	}

	curr := &Node[T]{
		data: data,
		next: nil,
	}
	temp := list.head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = curr
}

// Length of list - O(N)
func (list *LinkedList[T]) Len() int {
	temp := list.head
	count := 0
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

// Traversal print
func (list *LinkedList[T]) Print() {
	curr := list.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

// Front se delete karna - O(1)
func (list *LinkedList[T]) PopFront() {
	if list.head == nil {
		fmt.Println("List is Empty")
		return
	}
	list.head = list.head.next
}

// End se delete karna - O(N)
func (list *LinkedList[T]) PopBack() {
	// FIX 1: Empty list case
	if list.head == nil {
		fmt.Println("List is Empty")
		return
	}

	// FIX 2: Single node case
	if list.head.next == nil {
		list.head = nil
		return
	}

	// FIX 3: Multi-node traversal
	temp := list.head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
}

// Recursive reverse helper
func reverseListHelper[T any](head *Node[T]) *Node[T] {
	if head == nil || head.next == nil {
		return head
	}

	newHead := reverseListHelper(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

// Reverse wrapper
func (list *LinkedList[T]) ReverseList() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	list.head = reverseListHelper(list.head)
}

func main() {
	fmt.Println("LinkedList Implementation in Go")
	list := LinkedList[int]{}

	list.InsertAtBeginning(10)
	list.InsertAtBeginning(20)
	list.InsertAtBeginning(30)

	list.InsertAtEnd(100)
	list.InsertAtEnd(200)
	list.InsertAtEnd(300)
	list.Print() // 30 20 10 100 200 300

	fmt.Println("Length of List:", list.Len())

	list.PopFront()
	list.Print() // 20 10 100 200 300

	list.PopBack()
	list.Print() // 20 10 100 200

	list.ReverseList()
	list.Print() // 200 100 10 20
}
