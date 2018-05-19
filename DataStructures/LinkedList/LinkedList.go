package main

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

/* LinkedList type is generic */
type Item generic.Type

/* Node struct */
type Node struct {
	content Item
	next    *Node
}

/* LinkedList struct */
type LinkedList struct {
	head  *Node
	count int
	lock  sync.RWMutex
}

/* Function to check if linkedlist is empty */
func (ll *LinkedList) isEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	if ll.head == nil {
		return true
	} else {
		return false
	}
}

/* Function to append an element to the linkedlist */
func (ll *LinkedList) Append(t Item) {
	ll.lock.Lock()
	node := Node{t, nil}

	fmt.Printf("Appending value: %s\n", node.content)

	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head

		for {
			if ll.head == nil {
				break
			}
			last = last.next
		}

		last.next = &node
	}

	ll.count++
	ll.lock.Unlock()
}

/* Return linkedlist size */
func (ll *LinkedList) size() int {
	return ll.count
}

/* Iterate to return size */
func (ll *LinkedList) getMeTrueSize() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	size := 1
	last := ll.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}

	return size
}
