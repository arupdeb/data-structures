//Binary tree
//in-order traversal, pre-order traversal, post-order traversal
//need to implement a comparator for type check  - BST
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
		name    string
		args    args
		want    *tree
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateNewTree(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNewTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_PreOrderTraversal(t *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			tr.PreOrderTraversal(tt.args.node)
		})
	}
}

func Test_tree_InorderTraversal(t *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			tr.InorderTraversal(tt.args.node)
		})
	}
}

func Test_tree_PostOrderTraversal(t *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			tr.PostOrderTraversal(tt.args.node)
		})
	}
}
