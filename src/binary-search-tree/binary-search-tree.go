package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	key  int
}

type Tree struct {
	root *Node
}

// insertion is an O(log n) operation provided that the tree is moderately balanced.
func (T *Tree) Insert(key int) {
	if T.root == nil {
		newItem := &Node {
			nil,
			nil,
			key,
		}
		T.root = newItem
		return
	}
	T.InsertNode(T.root, key)
}

// The insertion algorithm is split for a good reason. The first algorithm (non- recursive)
// checks a very core base case - whether or not the tree is empty. If the tree is empty then
// we simply create our root node and finish. In all other cases we invoke the recursive InsertNode
// algorithm which simply guides us to the first appropriate place in the tree to put value.
func (T *Tree) InsertNode(current *Node, key int) {
	newItem := &Node {
		nil,
		nil,
		key,
	}

	if key < current.key {
		if current.left == nil {
			current.left = newItem
		} else {
			T.InsertNode(current.left, key)
		}
	} else {
		if current.right == nil {
			current.right = newItem
		} else {
			T.InsertNode(current.right, key)
		}
	}
}

// When using the pre order algorithm, you visit the root first,
// then traverse the left subtree and finally traverse the right subtree
func (T *Tree) PreOrder(current *Node) {
	if current != nil {
		T.PreOrder(current.left)
		T.PreOrder(current.right)
	}
}

// Search
func (T *Tree) Contains(current *Node, key int) bool {
	if current == nil {
		return false
	}

	if current.key == key {
		return true
	}
	if key < current.key {
		return T.Contains(current.left, key)
	}
	if key > current.key {
		return T.Contains(current.right, key)
	}

	return false
}

func (T *Tree) FindNode(current *Node, key int) *Node {
	if current == nil {
		return nil
	}

	if current.key == key {
		return current
	}
	if key < current.key {
		return T.FindNode(current.left, key)
	}
	if key > current.key {
		return T.FindNode(current.right, key)
	}

	return nil
}

func (T *Tree) FindParent(current *Node, key int) *Node {
	if current.key == key {
		return nil
	}
	if key < current.key {
		if current.left == nil {
			return nil
		} else if current.left.key == key {
			return current
		} else {
			return T.FindParent(current.left, key)
		}
	} else {
		if current.right == nil {
			return nil
		} else if current.right.key == key {
			return current
		} else {
			return T.FindParent(current.right, key)
		}
	}


	return nil
}

// Deletion
// Cases to consider:
// 1. the value to remove is a leaf node; or
// 2. the value to remove has a right subtree, but no left subtree; or
// 3. the value to remove has a left subtree, but no right subtree; or
// 4. the value to remove has both a left and right subtree in which case we promote the largest value in the left subtree.
func (T *Tree) Delete(key int) bool {
	nodeToRemove := T.FindNode(T.root, key)
	parent := T.FindParent(T.root, key)

	if nodeToRemove == nil {
		return false
	}

	if nodeToRemove.left == nil && nodeToRemove.right == nil {
		if nodeToRemove.key < parent.key {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if nodeToRemove.left == nil && nodeToRemove.right != nil {
		if nodeToRemove.key < parent.key {
			parent.left = nodeToRemove.right
		} else {
			parent.right = nodeToRemove.right
		}
	} else if nodeToRemove.left != nil && nodeToRemove.right == nil {
		if nodeToRemove.key < parent.key {
			parent.left = nodeToRemove.left
		} else {
			parent.right = nodeToRemove.left
		}
	} else {
		largestValue := nodeToRemove.left

		for largestValue.right != nil {
			// find the largest value in the left subtree of nodeToRemove
			largestValue = largestValue.right
		}

		// set the parent's Right pointer of largestValue to nil
		largestValueParent := T.FindParent(T.root, largestValue.key)
		largestValueParent.right = nil
		nodeToRemove.key = largestValue.key
	}

	return true
}
