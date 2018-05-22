package main

import (
	"testing"
)

var myList LinkedList

func TestAppend(t *testing.T) {
	if !myList.isEmpty() {
		t.Errorf("List should be empty")
	}

	myList.Append("Alexander")

	if myList.isEmpty() {
		t.Errorf("List should not be empty")
	}

	myList.Append("Aaron")
	myList.Append("Marquis")

	ListSize := myList.getMeTrueSize()
	if ListSize != 3 {
		t.Errorf("List size is not 3")
	}

	/* Iterate my linkedlist */
	myList.IterateList()
}

/* Testing IndexOf method */
func TestIndexOf(t *testing.T) {

	if i := myList.IndexOf("Alexander"); i != 0 {
		t.Errorf("Error to get index of first element: Alexander. Expected 0 but got %d", i)
	}

	if i := myList.IndexOf("Aaron"); i != 1 {
		t.Errorf("Error to get index of first element: Alexander. Expected 1 but got %d", i)
	}

	if i := myList.IndexOf("Marquis"); i != 2 {
		t.Errorf("Error to get index of first element: Alexander. Expected 2 but got %d", i)
	}

	if i := myList.IndexOf("Angelica"); i != -1 {
		t.Errorf("Error to get index of first element: Alexander. Expected -1 but got %d", i)
	}
}

/* Testing InsertAt method */
func (ll *LinkedList) TestInsertAt(t *testing.T) {
	err := myList.InsertAt("Eliza", 2)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err)
	}

	myList.Append("John")
	if i := myList.getMeTrueSize(); i != 5 {
		t.Errorf("Length of list is incorrect. Expected 5 but got %d\n", i)
	}

	myList.InsertAt("Lin", 0)
	if head := myList.head.content; head != "Lin" {
		t.Errorf("Value of head should be Lin but got %s", head)
	}

	/* Iterate list after insert test*/
	myList.IterateList()
}
