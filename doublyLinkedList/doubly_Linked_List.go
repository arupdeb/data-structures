package doublyLinkedList

import (
	"errors"
	"log"
)

type node struct {
	prev *node
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
		prev: nil,
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
	newNode.prev = ll.tail
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
		prev: nil,
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
		prev: ll.tail,
		data: element,
		next: nil,
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
	n := ll.head.next
	n.prev = nil 
	ll.head = n
	ll.size--
	return removedData, nil
}

//RemoveLast removes the last element of the linked list : complexity O(1)
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
	} else {
		removedData = ll.tail.data
		n := ll.tail.prev
		n.next = nil
		ll.tail = n
		ll.size--
	}
	return removedData, nil
}

func (ll *Links) Remove(n *node) (interface{}, error) {
	var removedData interface{}
	var err error = nil
	if ll.Size() == 0 {
		return nil, errors.New("empty List: cannot remove element")
	} else if n.next == nil {
		removedData, err = ll.RemoveLast()
		return removedData, err
	} else if n.prev == nil {
		removedData, err = ll.RemoveFirst()
		return removedData, err
	} else {
		for i := ll.head.next; i.next != nil; i = i.next {
			if i.next.next == n.next && i.next.prev == n.prev {
				removedData = i.next.data
				n.next.prev = i
				i.next = n.next
				ll.size--
				break
			}
		}
	}
	return removedData, err
}

//RemoveAt removes the node at a given index: not real index: Complexity O(n)
func (ll *Links) RemoveAt(index int) (interface{}, error) {
	var removedData interface{}
	var err error = nil

	if index >= ll.Size() {
		return nil, errors.New("index out of bounds")
	} else if index == 0 {
		removedData, err = ll.RemoveFirst()
		return removedData, err
	} else if index == ll.Size()-1 {
		removedData, err = ll.RemoveLast()
		return removedData, err
	}
	n := ll.head.next
	for i := 1; i < ll.Size()-1; i++ {
		if i == index {
			removedData = n.data
			n.prev.next = n.next
			n.next.prev = n.prev
			ll.size--
			break
		}
		n = n.next
	}

	return removedData, err
}

//RemoveElement removes a node matching the data passed : complexity O(n)
func (ll *Links) RemoveElement(element interface{}) (interface{}, error) {
	var removedData interface{} = nil
	var err error = nil
	for n := ll.head; n != nil; n = n.next {
		if n.data == element {
			if n.prev == nil {
				removedData, err = ll.RemoveFirst()
				break
			} else if n.next == nil {
				removedData, err = ll.RemoveLast()
				break
			} else {
				removedData, err =  ll.Remove(n)
				break
			}
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
