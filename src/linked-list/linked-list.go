package linkedList

import (
	"fmt"
)

type Node struct {
	next *Node
	key  int
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key int) {
	newItem := &Node{
		next: L.head,
		key:  key,
	}

	// update list head
	L.head = newItem

	// update list tail
	i := L.head
	for i.next != nil {
		i = i.next
	}
	L.tail = i
}

func (L *List) Remove(key int) {
	// case list is empty
	if L.head == nil {
		return
	}

	i := L.head
	if i.key == key {
		if i.next == nil {
			// case value is head and list has head only
			L.head = nil
			L.tail = nil
		} else {
			// case value is head and list contains more items
			L.head = L.head.next
		}
	} else {
		// search and stop the index to the previous item
		// before the one to search
		for i.next != nil && i.next.key != key {
			i = i.next
		}
		if i.next != nil {
			if i.next == L.tail {
				// case item is tail
				i.next = nil
				L.tail = i
			} else {
				// case item is found and it's not tail
				i.next = i.next.next
			}
		}
	}
}

func (L *List) Traverse() {
	i := L.head
	for i != nil {
		fmt.Printf("%+v\n", i.key)
		i = i.next
	}
}

func (L *List) ReverseTraversal() {
	if L.tail != nil {
		curr := L.tail
		for curr != L.head {
			// find previous
			prev := L.head
			for prev.next != curr {
				prev = prev.next
			}
			fmt.Printf("%+v\n", curr.key)
			curr = prev
		}
		fmt.Printf("%+v\n", curr.key)
	}
}
