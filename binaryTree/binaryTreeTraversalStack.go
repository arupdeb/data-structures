package binaryTree

import (
	"log"

	"github.com/arupdeb/data-structures/stacks"
)

//InorderTraversalStack : uses stack to traverse the binary tree using in-order traversal- DFS
func (t *tree) InorderTraversalStack(root *Node) {
	s := stacks.CreateNewStack()

	for i := root; i != nil || !s.IsEmpty(); {

		for j := i; j != nil; {
			s.Push(j)
			j = j.left
		}
		//j is nil after visiting all the left nodes

		// get the topmost entry and pop the value i.e visisted
		top, err := s.Pop()
		if err != nil {
			log.Println(err.Error())
		}

		// store the data to i for visiting
		i = top.(*Node)
		log.Println(i.data)

		// visit right subtree
		i = i.right
	}
}

//PreOrderTraversalStack : uses stack to traverse the binary tree using pre-order traversal- DFS
func (t *tree) PreOrderTraversalStack(root *Node) {
	s := stacks.CreateNewStack()

	for i := root; i != nil || !s.IsEmpty(); {

		for j := i; j != nil; {
			log.Println(j.data)
			s.Push(j)
			j = j.left
		}
		//j is nil after visiting all the left nodes

		// get the topmost entry and pop the value i.e visisted
		top, err := s.Pop()
		if err != nil {
			log.Println(err.Error())
		}

		// store the data to i for visiting
		i = top.(*Node)

		// visit right subtree
		i = i.right
	}
}

//PreOrderTraversalStack : uses stack to traverse the binary tree using pre-order traversal- DFS
func (t *tree) PostOrderTraversalStack(root *Node) {
	s := stacks.CreateNewStack()

	for i := root; i != nil || !s.IsEmpty(); {

		for j := i; j != nil; {
			s.Push(j)
			j = j.left
		}
		//j is nil after visiting all the left nodes

		// get the topmost entry and peek the value
		top, err := s.Peek()
		if err != nil {
			log.Println(err.Error())
		}

		// store the data to i for visiting
		i = top.(*Node)

		//Print the right node value
		if i.right != nil {
			log.Println("i.right.data : ",i.right.data)
		}
		//Print the Root node value
		log.Println("i.data : ",i.data)
		s.Pop() // visited
		// visit right subtree
		i = i.right

	}
}
