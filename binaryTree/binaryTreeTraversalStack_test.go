package binaryTree

import (
	"testing"
)

func Test_tree_InorderTraversalStack(t *testing.T) {

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
		{
			name: "Traverse only root node and left node and right node and more",
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
				t.root.right.right = &Node{
					data:  444,
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
			tr.InorderTraversalStack(tr.root)
		})
	}
}

func Test_tree_PreOrderTraversalStack(t *testing.T) {

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
				t.root.right.right = &Node{
					data:  612,
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
			tr.PreOrderTraversalStack(tr.root)
		})
	}
}

func Test_tree_PreOrderTraversalStack2(t *testing.T) {

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
				t.root.right.right = &Node{
					data:  512,
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
			tr.PreOrderTraversalStack2(tr.root)
		})
	}
}

func Test_tree_PostOrderTraversalStack(t *testing.T) {
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
		{
			name: "Traverse only root node and left node and right node and more",
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
				t.root.right.right = &Node{ 
					data:  624,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
			{
				name: "Traverse the tree from GFG example",
				beforeTest: func() *tree {
					t := CreateNewTree(1)
					t.root.left = &Node{
						data:  2,
						left:  nil,
						right: nil,
					}
					t.root.left.left = &Node{
						data:  4,
						left:  nil,
						right: nil,
					}
					t.root.left.right = &Node{
						data:  5,
						left:  nil,
						right: nil,
					}
					t.root.right = &Node{
						data:  3,
						left:  nil,
						right: nil,
					}
					t.root.right.left = &Node{ 
						data:  6,
						left:  nil,
						right: nil,
					}
					t.root.right.right = &Node{ 
						data:  7,
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
			tr.PostOrderTraversalStack(tr.root)
		})
	}
}

func Test_tree_PostOrderTraversalStackHashing(t *testing.T) {
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
		{
			name: "Traverse only root node and left node and right node and more",
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
				t.root.right.right = &Node{ 
					data:  624,
					left:  nil,
					right: nil,
				}
				return t
			},
		},
			{
				name: "Traverse the tree from GFG example",
				beforeTest: func() *tree {
					t := CreateNewTree(1)
					t.root.left = &Node{
						data:  2,
						left:  nil,
						right: nil,
					}
					t.root.left.left = &Node{
						data:  4,
						left:  nil,
						right: nil,
					}
					t.root.left.right = &Node{
						data:  5,
						left:  nil,
						right: nil,
					}
					t.root.right = &Node{
						data:  3,
						left:  nil,
						right: nil,
					}
					t.root.right.left = &Node{ 
						data:  6,
						left:  nil,
						right: nil,
					}
					t.root.right.right = &Node{ 
						data:  7,
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
			tr.PostOrderTraversalStackHashing(tr.root)
		})
	}
}
