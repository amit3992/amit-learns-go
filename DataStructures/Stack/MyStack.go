package Stack

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

/* Generic Item to populate the stack */
type Item generic.Type

/* Declare stack struct */
type MyStack struct {
	items []Item
	lock  sync.RWMutex
}

/* Func to create new stack */
func (s *MyStack) CreateStack() *MyStack {
	s.items = []Item{}
	return s
}

/* Push an element into the stack */
func (s *MyStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, t)
}

/* Pop an element from the stack */
func (s *MyStack) Pop() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()

	size := len(s.items)
	item := s.items[size-1]
	s.items = s.items[0 : size-1]

	return &item
}

/* Check if stack is empty */
func (s *MyStack) IsEmpty() bool {
	return len(s.items) == 0
}
