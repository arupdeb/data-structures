package linkedList_test

import (
	"testing"
	"github.com/arupdeb/data-structures/linkedList"
)

func TestCreateNewList_empty(t *testing.T){
	list := linkedList.CreateNewList()
	if result:=list.Size(); result != 0 {
		t.Errorf("Expected %v got %v", 0, result)
	}
}

func TestCreateNewList_intialData(t *testing.T){
	list := linkedList.CreateNewList("a", "b", "c", "d")
	if result:=list.Size(); result != 4 {
		t.Errorf("Expected %v got %v", 4, result)
	}
}