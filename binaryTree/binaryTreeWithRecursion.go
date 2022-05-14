//Binary tree
//in-order traversal, pre-order traversal, post-order traversal
//need to implement a comparator for type check  - BST
package binaryTree

import (
	"errors"
	"log"
)

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type tree struct {
	root *Node
	size int //total number of nodes in a tree
}

func CreateNewTree(value interface{}) (*tree, error) {
	n := &Node{
		data:  value,
		left:  nil,
		right: nil,
	}
	if value != nil {
		return &tree{
			root: n,
			size: 1,
		}, nil
	}
	return nil, errors.New("cannot create new tree: no data for root node")
}


func (t *tree) PreOrderTraversal(node *Node) {
	if node != nil {
		log.Println(node.data)
		t.PreOrderTraversal(node.left)
		t.PreOrderTraversal(node.right)
	}
}

func (t *tree) InorderTraversal(node *Node) {
	if node != nil {
		t.InorderTraversal(node.left)
		log.Println(node.data)
		t.InorderTraversal(node.right)
	}
}

func (t *tree) PostOrderTraversal(node *Node) {
	if node != nil {
		t.PostOrderTraversal(node.left)
		t.PostOrderTraversal(node.right)
		log.Println(node.data)
	}
}
