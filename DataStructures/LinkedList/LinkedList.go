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

/* Util methods */
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

/* Return head of linkedList */
func (ll *LinkedList) getListHead() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	return ll.head
}

/* Iterate list */
func (ll *LinkedList) IterateList() {

	fmt.Println("Iterating list:")
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	current := ll.head
	index := 0
	for {
		if current == nil {
			break
		}
		index++
		fmt.Print(current.content)
		fmt.Print(" ")
		current = current.next
	}

	fmt.Println()
}

/* ============================ FUNCTIONAL METHODS ================================ */

/* Function to append an element to the linkedlist */
func (ll *LinkedList) Append(t Item) {
	ll.lock.Lock()
	node := Node{t, nil}

	//fmt.Printf("Appending value: %s\n", node.content)

	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head

		for {
			if last.next == nil {
				break
			}
			last = last.next
		}

		last.next = &node
	}

	ll.count++
	ll.lock.Unlock()
}

/* Function to insert an element at a particular position in the linkedlist */
func (ll *LinkedList) InsertAt(t Item, pos int) error {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	if ll.head == nil || pos < 0 || ll.count < pos {
		return fmt.Errorf("Index out of bounds")
	}

	newNode := Node{t, nil}
	current := ll.head
	index := 0

	if pos == 0 {
		newNode.next = ll.head
		ll.head = &newNode
		ll.count++
		return nil
	}

	for index < pos-2 {
		index++
		current = current.next
	}

	newNode.next = current.next
	current.next = &newNode
	ll.count++
	return nil
}

/* Function which returns the index of an element in the linkedlist */
func (ll *LinkedList) IndexOf(t Item) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	if ll.head == nil {
		return 0
	}

	node := ll.head
	j := 0

	for {
		if node.content == t {
			return j
		}

		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}
