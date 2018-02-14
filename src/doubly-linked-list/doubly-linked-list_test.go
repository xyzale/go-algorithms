package main

import "testing"

func TestList_Insert(t *testing.T) {

	myList := List{}
	myList.Insert(1)
	if myList.head.key != 1 {
		t.Error("Wrong head")
	}
	myList.Insert(2)
	if myList.head.key != 1 {
		t.Error("Wrong head")
	}
	if myList.tail.key != 2 {
		t.Error("Wrong tail")
	}
}
