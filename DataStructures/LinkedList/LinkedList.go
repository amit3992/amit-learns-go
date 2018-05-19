package LinkedList

import (
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
	return ll.head == nil
}

/* Function to append an element to the linkedlist */
func (ll *LinkedList) Append(t Item) {
	ll.lock.Lock()
	node := Node{t, nil}

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
