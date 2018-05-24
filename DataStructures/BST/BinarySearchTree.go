package binarysearchtree

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

/* Item type for the BST */
type Item generic.Type

/* Node struct */
type Node struct {
	key   int
	value Item
	left  *Node
	right *Node
}

/* BST struct */
type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

/* Insert item t in the BST */
func (bst *BinarySearchTree) Insert(key int, value Item) {

	bst.lock.Lock()
	defer bst.lock.Unlock()

	/* Create new node */
	n := &Node{key, value, nil, nil}

	if bst.root == nil {
		bst.root = n
	} else {
		insertNodeUtil(bst.root, n)
	}
}

/* Internal util method to insert node into the BST */
func insertNodeUtil(node, newNode *Node) {

	if newNode.key < node.key {

		if node.left == nil {
			node.left = newNode
		} else {
			insertNodeUtil(node.left, newNode)
		}

	} else {

		if node.right == nil {
			node.right = newNode
		} else {
			insertNodeUtil(node.right, newNode)
		}
	}
}

/* Remove a node from BST */
func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	remove(bst.root, key)
}

/* Internal method to remove node from BST */
func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}

	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	/* Node.key == key */
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	leftMostNode := node.right
	/* Find smallest value in the right sub-tree of the node */
	for {
		if leftMostNode != nil && leftMostNode.left != nil {
			leftMostNode = leftMostNode.left
		} else {
			break
		}
	}

	node.key, node.value = leftMostNode.key, leftMostNode.value
	node.right = remove(node.right, node.key)
	return node
}

/* Search for an Item in the BST */
func (bst *BinarySearchTree) Search(key int) bool {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return search(bst.root, key)
}

/* Internal method to search in the BST */
func search(node *Node, key int) bool {

	if node == nil {
		return false
	}

	if key < node.key {
		return search(node.left, key)
	} else if key > node.key {
		return search(node.right, key)
	}

	return true

}

/* Min returns the Node with the minimum key value */
func (bst *BinarySearchTree) Min() *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := bst.root

	if node == nil {
		return nil
	}

	for {
		if node.left == nil {
			return node
		} else {
			node = node.left
		}
	}
}

/* Min returns the Node with the minimum key value */
func (bst *BinarySearchTree) Max() *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	node := bst.root

	if node == nil {
		return nil
	}

	for {
		if node.rigt == nil {
			return node
		} else {
			node = node.right
		}
	}
}

/* InOrderTraverse traverses all the nodes in order */
func (bst *BinarySearchTree) InOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	inOrderTraverseUtil(bst.root, f)
}

/* Internal util method to traverse in-order */
func inOrderTraverseUtil(n *Node, f func(Item)) {
	if n != nil {
		inOrderTraverseUtil(n.left, f)
		f(n.value)
		inOrderTraverseUtil(n.right, f)
	}
}

/* PreOrderTraverse traverses all the nodes in DFS preorder */
func (bst *BinarySearchTree) PreOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverseUtil(bst.root, f)
}

/* Internal util method to traverse pre-order */
func preOrderTraverseUtil(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)
		preOrderTraverseUtil(n.left, f)
		preOrderTraverseUtil(n.right, f)
	}
}

/* PostOrderTraverse traverses all the nodes in DFS postorder */
func (bst *BinarySearchTree) PostOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverseUtil(bst.root, f)
}

/* Internal util method to traverse post-order */
func postOrderTraverseUtil(n *Node, f func(Item)) {
	if n != nil {
		postOrderTraverseUtil(n.left, f)
		postOrderTraverseUtil(n.right, f)
		f(n.value)
	}
}

/* Print visual representation of the BST */
func (bst *BinarySearchTree) ToString() {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

/* Internal util method to print the BST */
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.left, level)
	}
}
