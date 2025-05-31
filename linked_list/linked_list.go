package main

import (
	"fmt"
)

type Node struct {
	next *Node
	num  int
}

type LinkedList struct {
	head   *Node
	length int
}

func (linkedlist *LinkedList) Append(value int) {
	if linkedlist.head == nil {
		linkedlist.head = &Node{num: value}
	} else {
		ptr := linkedlist.head

		for ptr.next != nil {
			ptr = ptr.next
		}

		ptr.next = &Node{num: value}
	}
	linkedlist.length++
}

func (linkedlist *LinkedList) PrintLinkedList() {
	ptr := linkedlist.head
	for ptr != nil {
		fmt.Printf("%v", ptr.num)
		ptr = ptr.next
	}
}

func (linkedlist *LinkedList) InsertAt(value int, pos int) {
	if pos == 0 {
		newNode := &Node{num: value, next: linkedlist.head}
		linkedlist.head = newNode
	} else {
		ptr := linkedlist.head
		i := 0
		for ptr != nil && i < pos-1 {
			ptr = ptr.next
			i++
		}
		if ptr == nil {
			return
		}

		newNode := &Node{num: value, next: ptr.next}
		ptr.next = newNode

	}
	linkedlist.length++
}

func (linkedlist *LinkedList) DeletebyIndex(pos int) {
	if linkedlist.head == nil || pos >= linkedlist.length {
		return
	}
	if pos == 0 {
		ptr := linkedlist.head
		linkedlist.head = ptr.next
		linkedlist.length--
		return
	} else {
		ptr := linkedlist.head
		i := 0

		for ptr != nil && i < pos-1 {
			ptr = ptr.next
			i++
		}

		if ptr == nil || ptr.next == nil {
			return // Nothing to delete
		}

		ptr.next = ptr.next.next
	}
	linkedlist.length--
}

func (linkedlist *LinkedList) Deletebyvalue(value int) {
	getIndex := linkedlist.SearchbyValue(value)

	if getIndex != -1 {
		linkedlist.DeletebyIndex(getIndex)
	}
}

func (linkedlist *LinkedList) SearchbyValue(value int) int {
	if linkedlist.head == nil {
		return -1
	}
	i := 0
	ptr := linkedlist.head

	for ptr != nil {
		if ptr.num == value {
			return i
		} else {
			ptr = ptr.next
			i++
		}
	}
	return -1
}

func main() {
	ll := &LinkedList{}

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	fmt.Print("After appending 10, 20, 30: ")
	ll.PrintLinkedList()
	fmt.Println()

	ll.InsertAt(15, 1)
	fmt.Print("After inserting 15 at index 1: ")
	ll.PrintLinkedList()
	fmt.Println()

	index := ll.SearchbyValue(20)
	fmt.Printf("Index of 20: %d\n", index)

	ll.DeletebyIndex(1)
	fmt.Print("After deleting index 1: ")
	ll.PrintLinkedList()
	fmt.Println()

	ll.Deletebyvalue(30)
	fmt.Print("After deleting value 30: ")
	ll.PrintLinkedList()
	fmt.Println()

	ll.DeletebyIndex(0)
	fmt.Print("After deleting head: ")
	ll.PrintLinkedList()
	fmt.Println()

	fmt.Print("Final list: ")
	ll.PrintLinkedList()
	fmt.Println()
}
