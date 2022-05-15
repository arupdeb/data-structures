package binaryTree

import (
	"testing"
)

func Test_tree_PreOrderTraversalRecursive(t *testing.T) {
	tests := []struct {
		name       string
		beforeTest func() *tree
	}{
		{
			name: "Traverse only root node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				return t
			},
		},
		{
			name: "Traverse only root node and left node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				t.root.left = &Node{
					data:  456,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
		{
			name: "Traverse only root node and left node and right node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				t.root.left = &Node{
					data:  456,
					left:  nil,
					right: nil,
				}
				t.root.right = &Node{
					data:  428,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.beforeTest()
			tr.PreOrderTraversalRecursive(tr.root)
		})
	}
}

func Test_tree_InorderTraversal(t *testing.T) {
	tests := []struct {
		name       string
		beforeTest func() *tree
	}{
		{
			name: "Traverse only root node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				return t
			},
		},
		{
			name: "Traverse only root node and left node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				t.root.left = &Node{
					data:  456,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
		{
			name: "Traverse only root node and left node and right node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				t.root.left = &Node{
					data:  456,
					left:  nil,
					right: nil,
				}
				t.root.right = &Node{
					data:  428,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.beforeTest()
			tr.InorderTraversalRecursive(tr.root)
		})
	}
}

func Test_tree_PostOrderTraversal(t *testing.T) {
	tests := []struct {
		name       string
		beforeTest func() *tree
	}{
		{
			name: "Traverse only root node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				return t
			},
		},
		{
			name: "Traverse only root node and left node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				t.root.left = &Node{
					data:  456,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
		{
			name: "Traverse only root node and left node and right node",
			beforeTest: func() *tree {
				t := CreateNewTree(123)
				t.root.left = &Node{
					data:  456,
					left:  nil,
					right: nil,
				}
				t.root.right = &Node{
					data:  428,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.beforeTest()
			tr.PostOrderTraversalRecursive(tr.root)
		})
	}
}
