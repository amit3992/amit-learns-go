package binarysearchtree

import (
	"fmt"
	"testing"
)

var bst BinarySearchTree

func BuildTree(bst *BinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

func TestInsert(t *testing.T) {
	BuildTree(&bst)
	bst.ToString()

	bst.Insert(11, "11")
	fmt.Println("After Inserting 11")
	bst.ToString()
}

func TestRemove(t *testing.T) {

	bst.Remove(11)
	fmt.Println("After Removing 11")
	bst.ToString()
}
