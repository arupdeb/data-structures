package queue

import (
	"errors"
	"log"
)

type queue struct {
	head *node
	tail *node
	size int
}
type node struct {
	data interface{}
	next *node
}

func CreateNewQueue(values ...interface{}) *queue {

	q := &queue{
		head: nil,
		tail: nil,
		size: 0,
	}
	if len(values) == 0 {
		return q
	}
	q.Enqueue(values...)
	return q
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Enqueue(values ...interface{}) {

	if len(values) == 0 {
		return
	}

	for i := 0; i < len(values); i++ {
		n := &node{
			data: values[i],
			next: nil,
		}

		if q.head == nil && q.tail == nil {
			q.head = n
			q.tail = n
			q.size++
			continue
		}

		q.tail.next = n
		q.tail = n
		q.size++
	}
}

func (q *queue) Dequeue() (interface{}, error) {
	if q.Size() == 0 {
		return nil, errors.New("cannot dequeue elememt: empty queue")
	}
	removedData := q.head.data
	q.head = q.head.next
	q.size--
	return removedData, nil
}

func (q *queue) Peek() (interface{}, error) {
	if q.Size() == 0 {
		return nil, errors.New("cannot peek elememt: empty queue")
	}
	return q.head.data, nil
}

func (q *queue) Remove(element interface{}) error {
	if q.Size() == 0 {
		return errors.New("cannot remove elememt: empty queue")
	}

	if q.head.data == element { // head element
		_, err := q.Dequeue()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	for i := q.head; i.next != nil; i = i.next {
		log.Println("i", i)

		if i.next.next == nil && i.next.data == element { //last element
			q.tail = i
			i.next = nil
			q.size--
			return nil
		} else if i.next.data == element {
			i.next = i.next.next
			q.size--
			return nil
		}

	}
	return errors.New("cannot find element to remove")
}

func (q *queue) PrintQueue() {
	if q.Size() == 0 {
		log.Println("cannot Print elememt: empty queue")
	}
	for i := q.head; i != nil; i = i.next {
		log.Println(i)
	}
}
