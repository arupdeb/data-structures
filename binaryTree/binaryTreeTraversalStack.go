package binaryTree

import (
	"log"

	"github.com/arupdeb/data-structures/stacks"
)

func (t *tree) InorderTraversalStack(root *Node) {
	s := stacks.CreateNewStack()

	for i := root; i != nil || !s.IsEmpty(); {

		for j := i; j != nil; {
			s.Push(j)
			j = j.left
		}

		// topmost entry of stack or head of the stack
		top, err := s.Peek()
		if err != nil {
			log.Println(err.Error())
		}

		// store the data to i for visiting
		i = top.(*Node)
		log.Println(i.data)

		// remove the top node visited
		s.Pop()

		// visit right subtree
		i = i.right
	}
}
