package LinkedList

import (
	"fmt"
	"sync"
	
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	content Item
	next *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

