//size, isEmpty, push
//pop, peek (return without removing)
//search (linear) O()n

package stacks

import (
	"errors"
	"log"
)

type Stack struct {
	head *node
	size int
}

type node struct {
	data interface{}
	next *node
}

func CreateNewStack(values ...interface{}) *Stack {
	s := &Stack{
		head: nil,
		size: 0,
	}
	if len(values) == 0 {
		return s
	}
	s.Push(values...)

	return s
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

//Push adds element to a stack not considering the max stack size and can accept nil value as data value of stack
func (s *Stack) Push(values ...interface{}) {

	if len(values) == 0 {
		return
	}
	for i := 0; i < len(values); i++ {
		n := &node{
			data: values[i],
			next: nil,
		}
		if s.head == nil {
			s.head = n
			s.size++
			continue
		}
		n.next = s.head
		s.head = n
		s.size++
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Size() == 0 {
		return nil, errors.New("cannot remove element: empty stack")
	}
	removedData := s.head.data
	s.head = s.head.next
	s.size--
	return removedData, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.Size() == 0 {
		return nil, errors.New("cannot read element: empty stack")
	}
	return s.head.data, nil
}

func (s *Stack) Contains(element interface{}) bool {
	if s.Size() == 0 {
		return false
	}
	for n := s.head; n != nil; n = n.next {
		if n.data == element {
			return true
		}
	}
	return false
}

func (s *Stack) PrintStack() {
	if s.Size() == 0 {
		log.Println("empty stack")
	}
	for n := s.head; n != nil; n = n.next {
		log.Println(n)
	}
}
