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
	myList.iterateList()
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
