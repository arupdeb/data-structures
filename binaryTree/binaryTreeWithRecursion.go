//Binary tree
//need to implement a comparator for type check  - BST
//TO-DO
// add_element(node), delete_element(), search_element(), bench_mark test
package binaryTree

import (
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

func CreateNewTree(value interface{}) *tree {
	if value == nil {
		return &tree{
			root: nil,
			size: 0,
		}
	}
	return &tree{
		root: &Node{
			data:  value,
			left:  nil,
			right: nil,
		},
		size: 1,
	}
}

func (t *tree) TreeSize() int {
	if t.root == nil {
		return 0
	}
	return t.size
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
