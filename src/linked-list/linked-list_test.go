package linkedList

import (
	"testing"
)

func Test_linkedList(t *testing.T) {
	myList := List{}
	myList.Insert(1)
	myList.Insert(2)
	myList.Insert(3)
	myList.Insert(4)
	myList.Insert(5)
	myList.Insert(6)

	if myList.head.key != 6 {
		t.Error("[error] Head is wrong")
	}
}