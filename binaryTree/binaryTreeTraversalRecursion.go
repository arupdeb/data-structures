package binaryTree

import (
	"log"
)

func (t *tree) PreOrderTraversalRecursive(node *Node) {
	if node != nil {
		log.Println(node.data)
		t.PreOrderTraversalRecursive(node.left)
		t.PreOrderTraversalRecursive(node.right)
	}
}

func (t *tree) InorderTraversalRecursive(node *Node) {
	if node != nil {
		t.InorderTraversalRecursive(node.left)
		log.Println(node.data)
		t.InorderTraversalRecursive(node.right)
	}
}

func (t *tree) PostOrderTraversalRecursive(node *Node) {
	if node != nil {
		t.PostOrderTraversalRecursive(node.left)
		t.PostOrderTraversalRecursive(node.right)
		log.Println(node.data)
	}
}

// add Benchmark after add() and delete() functions are available