package binaryTree

import (
	"reflect"
	"testing"
)

func TestCreateNewTree(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "create new Tree empty paramenter or nil",
			args: args{
				value: nil,
			},
			want: 0,
		},
		{
			name: "create new tree with a value",
			args: args{
				value: 45687,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateNewTree(tt.args.value)
			if !reflect.DeepEqual(got.TreeSize(), tt.want) {
				t.Errorf("CreateNewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_PreOrderTraversal(t *testing.T) {
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
					data: 456,
					left: nil,
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
					data: 456,
					left: nil,
					right: nil,
				}
				t.root.right = &Node{
					data: 428,
					left: nil,
					right: nil,
				}
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.beforeTest()
			tr.PreOrderTraversal(tr.root)
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
					data: 456,
					left: nil,
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
					data: 456,
					left: nil,
					right: nil,
				}
				t.root.right = &Node{
					data: 428,
					left: nil,
					right: nil,
				}
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.beforeTest()
			tr.PreOrderTraversal(tr.root)
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
					data: 456,
					left: nil,
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
					data: 456,
					left: nil,
					right: nil,
				}
				t.root.right = &Node{
					data: 428,
					left: nil,
					right: nil,
				}
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.beforeTest()
			tr.PostOrderTraversal(tr.root)
		})
	}
}
