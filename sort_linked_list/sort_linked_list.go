// Sort a linkedlist

package main

import "fmt"

type Node struct {
	next *Node
	num  int
}

type LinkedList struct {
	head *Node
}

func AddValues(linkedlist *LinkedList, value int) {
	if linkedlist.head == nil {
		linkedlist.head = &Node{num: value}
	} else {
		ptr := linkedlist.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = &Node{num: value}
	}
}

func PrintLinkedList(linkedlist *LinkedList) {
	ptr := linkedlist.head

	for ptr != nil {
		fmt.Printf("%v", ptr.num)
		ptr = ptr.next
	}
}

func SortLinkedList(linkedlist *LinkedList) *LinkedList {

}
