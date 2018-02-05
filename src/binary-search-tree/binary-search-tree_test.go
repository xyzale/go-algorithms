package main

import (
	"testing"
)

func Test_binarySearchTreeInsert(t *testing.T) {
	tree := Tree{}
	tree.Insert(10)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(23)
	tree.Insert(14)
	tree.Insert(38)
	tree.Insert(21)

	tree.PreOrder(tree.root)

	if !tree.Contains(tree.root, 21) {
		t.Error("21 not found")
	}
}

func Test_binarySearchTreeDelete(t *testing.T) {
	tree := Tree{}
	tree.Insert(10)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(23)
	tree.Insert(14)
	tree.Insert(38)
	tree.Insert(21)

	tree.PreOrder(tree.root)
	tree.Delete(14)

	if tree.Contains(tree.root, 14) {
		t.Error("14 not deleted")
	}
}
