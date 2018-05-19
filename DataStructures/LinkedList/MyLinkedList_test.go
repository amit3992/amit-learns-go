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
	/*
		myList.Append("Aaron")

			myList.Append("Marquis")

				ListSize := myList.getMeTrueSize()
				if ListSize != 3 {
					t.Errorf("List size is not 3")
				} */
}
