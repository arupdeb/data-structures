package binaryTree

import (
	"log"

	"github.com/arupdeb/data-structures/stacks"
)

//InorderTraversalStack : uses stack to traverse the binary tree using in-order traversal- DFS
func (t *tree) InorderTraversalStack(root *Node) {

	if root == nil {
		return
	}
	s := stacks.CreateNewStack()
	var curr *Node = root

	for curr != nil || !s.IsEmpty() {

		for curr != nil {
			s.Push(curr)
			curr = curr.left
		}
		//curr is nil after visiting all the left nodes
		// get the topmost entry and pop the value i.e visisted
		top, err := s.Pop()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		// store the data to i for visiting
		log.Println(top.(*Node).data)
		curr = top.(*Node)
		// visit right subtree
			curr = curr.right
		
	}
}

//PreOrderTraversalStack : uses stack to traverse the binary tree using pre-order traversal- DFS
func (t *tree) PreOrderTraversalStack(root *Node) {
	if root == nil {
		return
	}

	s := stacks.CreateNewStack()
	s.Push(root)

	for !s.IsEmpty() {

		//pop the first element if not empty
		headElement, err := s.Pop()
		if err != nil {
			log.Println(err.Error())
			break
		}
		data := headElement.(*Node)
		log.Println(data.data) // Print the root node

		if data.right != nil {
			s.Push(data.right)
		}
		if data.left != nil {
			s.Push(data.left)
		}
	}
}

//PreOrderTraversalStack2  2nd method: uses stack to traverse the binary tree using pre-order traversal- DFS
func (t *tree) PreOrderTraversalStack2(root *Node) {
	if root == nil {
		return
	}

	s := stacks.CreateNewStack()
	var curr *Node = root //initalize current variable

	for !s.IsEmpty() || curr != nil {

		//Print the left/root node first and push the right child to the stack
		//update the current node to the left node
		for curr != nil {
			log.Println(curr.data)

			if curr.right != nil {
				s.Push(curr.right)
			}
			curr = curr.left
		}
		//here curr will be nil after the loop completes
		//take out the topmost from stack i.e right child of the node printed above
		//set current to the right child / popped element 
		if !s.IsEmpty() {
			val, err := s.Pop()
			if err != nil {
				log.Println(err.Error())
				continue
			}
			curr = val.(*Node)
		}
	}
}

// PostOrderTraversalStack : uses stack to traverse the binary tree using post-order traversal- DFS
// func (t *tree) PostOrderTraversalStack(root *Node) {
// 	s := stacks.CreateNewStack()

// 	for i := root; i != nil || !s.IsEmpty(); {

// 		for j := i; j != nil; {
// 			s.Push(j)
// 			j = j.left
// 		}
// 		//j is nil after visiting all the left nodes

// 		// get the topmost entry and peek the value
// 		top, err := s.Peek()
// 		if err != nil {
// 			log.Println(err.Error())
// 		}

// 		// store the data to i for visiting
// 		i = top.(*Node)

// 		//Print the right node value
// 		if i.right != nil {
// 			log.Println("i.right.data : ", i.right.data)
// 		}
// 		//Print the Root node value
// 		log.Println("i.data : ", i.data)
// 		s.Pop() // visited
// 		// visit right subtree
// 		i = i.right

// 	}
// }
