package linkedList

import (
	"errors"
	"log"
)

type node struct {
	data interface{}
	next *node
}

type Links struct {
	head *node
	tail *node
	size int
}

func CreateNewList(values ...interface{}) *Links {
	ll := &Links{
		head: nil,
		tail: nil,
		size: 0,
	}
	if len(values) == 0 {
		return ll
	}
	for i := 0; i < len(values); i++ {
		ll.Add(values[i])
	}
	return ll
}

func (ll *Links) Add(element interface{}) {
	newNode := &node{
		data: element,
		next: nil,
	}
	if ll.size == 0 {
		ll.head = newNode
		ll.tail = newNode
		ll.size++
		return
	}
	ll.tail.next = newNode
	ll.tail = newNode
	ll.size++
}

func (ll *Links) Size() int {
	return ll.size
}

func (ll *Links) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

func (ll *Links) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *Links) AddFirst(element interface{}) {
	if ll.Size() == 0 {
		ll.Add(element)
		return
	}
	newNode := &node{
		data: element,
		next: ll.head,
	}
	ll.head = newNode
	ll.size++
}

func (ll *Links) AddLast(element interface{}) {
	if ll.Size() == 0 {
		ll.Add(element)
		return
	}
	newNode := &node{
		data: element,
		next: ll.tail,
	}
	ll.tail = newNode
	ll.size++
}

func (ll *Links) PeekFirst() (interface{}, error) {
	if ll.Size() == 0 {
		return nil, errors.New("empty List")
	}
	return ll.head.data, nil
}

func (ll *Links) PeekLast() (interface{}, error) {
	if ll.Size() == 0 {
		return nil, errors.New("empty List")
	}
	return ll.tail.data, nil
}

func (ll *Links) RemoveFirst() (interface{}, error) {
	var removedData interface{}
	if ll.Size() == 0 {
		return nil, errors.New("empty List: cannot remove element")
	} else if ll.Size() == 1 {
		removedData = ll.head.data
		ll.head = nil
		ll.tail = nil
		ll.size = 0
		return removedData, nil
	}
	removedData = ll.head.data
	ll.head = ll.head.next
	ll.size--
	return removedData, nil
}

//RemoveLast removes the last element of the linked list : complexity O(n)
func (ll *Links) RemoveLast() (interface{}, error) {
	var removedData interface{}
	if ll.Size() == 0 {
		return nil, errors.New("empty List: cannot remove element")
	} else if ll.Size() == 1 {
		removedData = ll.head.data
		ll.head = nil
		ll.tail = nil
		ll.size = 0
		return removedData, nil
	}
	var i *node
	for i = ll.head; i.next != nil; i = i.next {
		if i.next.next == nil {
			removedData = i.next.data
			i.next = nil
			ll.tail = i
			ll.size--
			break
		}
	}
	return removedData, nil
}

func (ll *Links) Remove(n *node) (interface{}, error) {
	var removedData interface{}
	var err error
	if ll.Size() == 0 {
		return nil, errors.New("empty List: cannot remove element")
	} else if n.next == nil {
		removedData, err = ll.RemoveLast()
		return removedData, err
	}
	for i := ll.head; i.next != nil; i = i.next {
		if i.next.next == n.next {
			removedData = i.next.data
			i.next = n.next
			ll.size--
			break
		}
	}
	return removedData, nil
}

//RemoveAt removes the node at a given index: not real index: Complexity O(n)
func (ll *Links) RemoveAt(index int) (interface{}, error) {
	var removedData interface{}
	var err error
	if index >= ll.Size() {
		return nil, errors.New("index out of bounds")
	} else if index == 0 {
		removedData, err = ll.RemoveFirst()
		return removedData, err
	} else if index == ll.Size()-1 {
		removedData, err = ll.RemoveLast()
		return removedData, err
	} else {
		var i int = 1
		for n := ll.head; n.next != nil; n = n.next {

			if i == index {
				removedData = n.next.data
				err = nil
				n.next = n.next.next
				ll.size--
				break
			}
			i++
		}
	}
	return removedData, err
}

//RemoveElement removes a node matching the data passed : complexity O(n)
func (ll *Links) RemoveElement(element interface{}) (interface{}, error) {
	var removedData interface{} = nil
	var err error = nil
	var i int = 0
	for n := ll.head; n != nil; n = n.next { // cannot be n.next != nil as nil pointer reference for empty array
		if i == 0 || i == ll.Size()-1 {
			if n.data == element && i == 0 {
				removedData, err = ll.RemoveFirst()
				break
			} else if n.data == element && i == ll.Size()-1 {
				removedData, err = ll.RemoveLast()
				break
			} else {
				i++
			}
		} else if n.next.data == element {
			removedData = n.next.data
			n.next = n.next.next
			ll.size--
			break
		} else {
			i++
		}
	}
	if removedData == nil {
		return nil, errors.New("element not found")
	}
	return removedData, err
}

//Contains returns true/false based on element found on the list : complexity O(n)
func (ll *Links) Contains(element interface{}) bool {
	for n := ll.head; n != nil; n = n.next {
		if n.data == element {
			return true
		}
	}
	return false
}

//PrintList : Prints the list in console
func (ll *Links) PrintList() {
	log.Println("Head: ", ll.head)
	log.Println("Tail: ", ll.tail)
	if ll.Size() > 0 {
		for n := ll.head; n != nil; n = n.next {
			log.Println(n)
		}
	}

}
