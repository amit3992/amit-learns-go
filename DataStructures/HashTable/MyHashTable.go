package HashTable

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

/* Generic key type */
type Key generic.Type

/* Generic value type */
type Value generic.Type

type MyHashTable struct {
	items map[int]Value
	lock  sync.RWMutex
}

/* Horner's method to generate hashcode of a string in O(n) */
func hash(k Key) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

/* Put item into the hashtable */
func (ht *MyHashTable) Put(k Key, v Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(k)

	if ht.items == nil {
		ht.items = make(map[int]Value)
	}

	ht.items[i] = v
}

/* Delete value with key -> k from the hashtable */
func (ht *MyHashTable) Delete(k Key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(k)
	delete(ht.items, i)
}

/* Get function to get the value from the hashtable */
func (ht *MyHashTable) Get(k Key) Value {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(k)
	return ht.items[i]
}

func (ht *MyHashTable) Size() int {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	return len(ht.items)
}
