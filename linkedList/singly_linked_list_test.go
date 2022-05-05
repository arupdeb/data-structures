package linkedList_test

import (
	"testing"

	"github.com/arupdeb/data-structures/linkedList"
)

func TestCreateNewList_empty(t *testing.T) {
	list := linkedList.CreateNewList()
	if result := list.Size(); result != 0 {
		t.Errorf("Expected %v got %v", 0, result)
	}
}

func TestCreateNewList_intialData(t *testing.T) {
	list := linkedList.CreateNewList("a", "b", "c", "d")
	if result := list.Size(); result != 4 {
		t.Errorf("Expected %v got %v", 4, result)
	}
}

func TestAdd_elements(t *testing.T) {
	list := linkedList.CreateNewList()
	list.Add("1000")
	if res := list.Size(); res != 1 {
		t.Errorf("Expected %v, got %v ", 1, res)
	}

	for i := 0; i < 100; i++ {
		list.Add(i)
	}
	//list.PrintList()
	if res := list.Size(); res != 101 {
		t.Errorf("expected %v, got %v", 101, res)
	}
}

func TestRemoveFirst(t *testing.T) {
	list := linkedList.CreateNewList()

	if _, err := list.RemoveFirst(); err == nil {
		t.Errorf("Expected %v, got %v", "empty List: cannot remove element", err.Error())
	}
	list.Add("1000")
	if res, err := list.RemoveFirst(); res == "1000" && err != nil {
		t.Errorf("Expected %v, got %v", "1000", res)
	}
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	//list.PrintList()
	if res, err := list.RemoveFirst(); res.(int) != 0 || err != nil {
		t.Errorf("Expected %v, got %v", "0", res)
	}
	//list.PrintList()
}

func TestRemoveLast(t *testing.T) {
	list := linkedList.CreateNewList()

	if _, err := list.RemoveLast(); err == nil {
		t.Errorf("Expected %v, got %v", "empty List: cannot remove element", err.Error())
	}
	list.Add("1000")                                                         // adding string to a linked list of Type interface
	if res, err := list.RemoveLast(); res.(string) == "1000" && err != nil { // so compare using string type cast
		t.Errorf("Expected %v, got %v", "1000", res)
	}
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	if res, err := list.RemoveLast(); res.(int) != 9 || err != nil { // so compare using int type cast in the same List
		t.Errorf("Expected %v, got %v", "9", res)
	}
}

func TestRemoveAt(t *testing.T) {
	list := linkedList.CreateNewList()

	if _, err := list.RemoveAt(0); err == nil {
		t.Errorf("Expected %v got %v", "index out of bounds", err)
	}

	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	if res, err := list.RemoveAt(0); res.(int) != 0 || err != nil { //imp: convert the interface type to int before compare
		t.Errorf("Expected %v got %v", "0", res)
	}
	if res, err := list.RemoveAt(5); res.(int) != 6 || err != nil {
		t.Errorf("Expected %v got %v", "6", res)
	}
	if res, err := list.RemoveAt(7); res.(int) != 9 || err != nil {
		t.Errorf("Expected %v got %v", 9, res)
	}
}
