package main

import "fmt"

type Node struct {
	previous *Node
	next *Node
	key  int
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key int) {
	newItem := &Node{
		previous: nil,
		next: nil,
		key:  key,
	}

	if L.head == nil {
		L.head = newItem
		L.tail = newItem
	} else {
		newItem.previous = L.tail
		L.tail.next = newItem
		L.tail = newItem
	}
}

func (L *List) Traverse() {
	i := L.head
	for i != nil {
		fmt.Printf("%+v\n", i.key)
		i = i.next
	}
}
