package doublyLinkedList_test

import (
	"testing"

	"github.com/arupdeb/data-structures/doublyLinkedList"
)

func TestCreateNewDoubleList_empty(t *testing.T) {
	list := doublyLinkedList.CreateNewList()
	if result := list.Size(); result != 0 {
		t.Errorf("Expected %v got %v", 0, result)
	}
}

func TestCreateNew_DoubleList_intialData(t *testing.T) {
	list := doublyLinkedList.CreateNewList("a", "b", "c", "d")
	// list.PrintList()
	if result := list.Size(); result != 4 {
		t.Errorf("Expected %v got %v", 4, result)
	}
}

func TestAdd_elements_DLL(t *testing.T) {
	list := doublyLinkedList.CreateNewList()
	list.Add("1000")
	if res := list.Size(); res != 1 {
		t.Errorf("Expected %v, got %v ", 1, res)
	}

	for i := 0; i < 100; i++ {
		list.Add(i)
	}
	// list.PrintList()
	if res := list.Size(); res != 101 {
		t.Errorf("expected %v, got %v", 101, res)
	}
}

func TestRemoveFirst_dll(t *testing.T) {
	list := doublyLinkedList.CreateNewList()

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
	// list.PrintList()
	if res, err := list.RemoveFirst(); res.(int) != 0 || err != nil {
		t.Errorf("Expected %v, got %v", "0", res)
	}
	// list.PrintList()
}

func TestRemoveLast_dll(t *testing.T) {
	list := doublyLinkedList.CreateNewList()

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
	// list.PrintList()
	if res, err := list.RemoveLast(); res.(int) != 9 || err != nil { // so compare using int type cast in the same List
		t.Errorf("Expected %v, got %v", "9", res)
	}
	// list.PrintList()
}


func TestRemoveAt_dll(t *testing.T) {
	list := doublyLinkedList.CreateNewList()

	if _, err := list.RemoveAt(0); err == nil {
		t.Errorf("Expected %v got %v", "index out of bounds", err)
	}

	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	// list.PrintList()
	if res, err := list.RemoveAt(0); res.(int) != 0 || err != nil { //imp: convert the interface type to int before compare
		t.Errorf("Expected %v got %v", "0", res)
	}
	// list.PrintList()

	if res, err := list.RemoveAt(5); res.(int) != 6 || err != nil {
		t.Errorf("Expected %v got %v", "6", res)
	}
	// list.PrintList()

	if res, err := list.RemoveAt(7); res.(int) != 9 || err != nil {
		t.Errorf("Expected %v got %v", 9, res)
	}
	// list.PrintList()

}

func TestRemoveElement_dll(t *testing.T) {
	list := doublyLinkedList.CreateNewList()
	if _, err := list.RemoveElement(32); err == nil {
		t.Errorf("Expected %v got %v", "element not found", err)
	}
	list.Add("abc") // adding combination of strings and integer in the same list
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	if res, err := list.RemoveElement(6); res != 6 || err != nil {
		t.Errorf("Expected %v, got %v", 6, res)
	}
	if res, err := list.RemoveElement("abc"); res != "abc" || err != nil {
		t.Errorf("Expected %v, got %v", "abc", res)
	}
	if res, err := list.RemoveElement(9); res != 9 || err != nil {
		t.Errorf("Expected %v, got %v", res, err)
	}
	if res, err := list.RemoveElement(4); res != 4 || err != nil {
		t.Errorf("Expected %v, got %v", res, err)
	}
}

func Test_Contains_dll(t *testing.T) {
	list := doublyLinkedList.CreateNewList()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	if res := list.Contains(0); res != true {
		t.Errorf("Expected %v, got %v", true, res)
	}
	if res := list.Contains(9); res != true {
		t.Errorf("Expected %v, got %v", true, res)
	}
	if res := list.Contains(5); res != true {
		t.Errorf("Expected %v, got %v", true, res)
	}
}

func Test_IsEmpty_dll(t *testing.T) {
	list := doublyLinkedList.CreateNewList()
	if res := list.IsEmpty(); res != true {
		t.Errorf("Expected %v, got %v", true, res)
	}
	list.Add("a")
	if res := list.IsEmpty(); res != false {
		t.Errorf("Expected %v, got %v", true, res)
	}
}

func BenchmarkAdd_dll(b *testing.B) {
	b.StopTimer()
	list := doublyLinkedList.CreateNewList()
	// for i := 0; i < 100; i++ {
	// 	list.Add(i)
	// }
	b.StartTimer()
	BenchMarkAdd_dll(b, list, 100)
}

func BenchMarkAdd_dll(b *testing.B, list *doublyLinkedList.Links, testSize int) {
	for i := 0; i <= b.N; i++ {
		for j := 0; j < testSize; j++ {
			list.Add(j)
		}
	}
}


func BenchmarkRemoveElement_dll(b *testing.B) {
	b.StopTimer()
	list := doublyLinkedList.CreateNewList()
	for i := 0; i < 100; i++ {
		list.Add(i)
	}
	b.StartTimer()
	BenchMarkele_dll(b, list, 100)
}

func BenchMarkele_dll(b *testing.B, list *doublyLinkedList.Links, testSize int) {
	for i := 0; i <= b.N; i++ {
		for j := testSize -1; j >=0 ; j-- {
			list.RemoveElement(j)
		}
	}
}