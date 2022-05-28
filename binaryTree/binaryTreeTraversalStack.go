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

func (t *tree) PostOrderTraversalStack(root *Node) {

	if root == nil {
		return
	}
	s := stacks.CreateNewStack()
	curr := root

	for !s.IsEmpty() || curr != nil {

		for curr != nil {
			if curr.right != nil {
				s.Push(curr.right) //push right node if exists: befor the root node
			}
			s.Push(curr) //push root node to stack after right child

			curr = curr.left
		}

		val, err := s.Pop() // pop the top element after current reaches null i.e root
		if err != nil {
			log.Println(err.Error())
		}
		curr = val.(*Node)

		if curr.right != nil && !s.IsEmpty() { // check if stack is empty and root node's right child is nil
			top, _ := s.Peek()
			// top of stack and root's right node are equal
			if top.(*Node).left == curr.right.left && top.(*Node).right == curr.right.right && top.(*Node).data == curr.right.data {
				root := curr       //store the root node in temp variable
				val, _ = s.Pop()   // pop the right child when equal to roots right child
				curr = val.(*Node) //set current node to the right child to traverse the tree
				s.Push(root)       //push the root back to stack
				continue			// continue to traverse the right tree without Printing 
			}

		}
		log.Println(curr.data) //Print the current node/ root node
		curr = nil

	}

}

// PostOrderTraversalStack : uses stack to traverse the binary tree using post-order traversal- DFS
func (t *tree) PostOrderTraversalStack(root *Node) {
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
 }
