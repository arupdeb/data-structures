//Binary tree
//need to implement a comparator for type check  - BST
//TO-DO
// add_element(node), delete_element(), search_element(), bench_mark test
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

func Test_tree_TreeSize(t *testing.T) {
	type fields struct {
		root *Node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tree{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := tr.TreeSize(); got != tt.want {
				t.Errorf("tree.TreeSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
