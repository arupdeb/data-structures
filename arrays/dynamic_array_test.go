package arrays_test

import (
	"testing"

	"github.com/arupdeb/data-structures/arrays"
)

func TestCreateNewEmptyArray(t *testing.T) {
	arr := arrays.CreateNewArray()
	if arraySize := arr.Size(); arraySize != 0 {
		t.Errorf("Expected: %v , Got %v", 0, arraySize)
	}
	if res := arr.IsEmpty(); res != true {
		t.Errorf("Expected: %v, Got %v", true, res)
	}
}

func TestCreateNewArray(t *testing.T) {
	arr := arrays.CreateNewArray("a", "b", "c")
	if arraySize := arr.Size(); arraySize != 3 {
		t.Errorf("Expected: %v , Got %v", 0, arraySize)
	}
	if res := arr.IsEmpty(); res != false {
		t.Errorf("Expected: %v, Got %v", true, res)
	}
}

func TestRemoveAt(t *testing.T) {
	arr := arrays.CreateNewArray("a", "b", "c", "d")
	if res, err := arr.RemoveAt(2); err != nil && res != "c" {
		t.Errorf("Expected: %v, Got %v", "c", res)
	}
	if res, err := arr.RemoveAt(2); err != nil && res != "d" {
		t.Errorf("Expected: %v, Got %v", "d", res)
	}
	if res, err := arr.RemoveAt(1); err != nil && res != "b" {
		t.Errorf("Expected: %v, Got %v", "d", res)
	}
	if res, err := arr.RemoveAt(0); err != nil && res != "b" {
		t.Errorf("Expected: %v, Got %v", "d", res)
	}
	if _, err := arr.RemoveAt(2); err == nil {
		t.Errorf("Expected: %v, Got %v", "array Index out of bounds", err.Error())
	}
}

func TestAdd(t *testing.T) {
	arr := arrays.CreateNewArray()
	arr.Add("a", "b", "c", "d")
	if size := arr.Size(); size != 4 {
		t.Errorf("Expected %v got %v", 4, size)
	}
	arr.Add("e")
	if size := arr.Size(); size != 5 {
		t.Errorf("Expected %v got %v", 5, size)
	}
}

func TestGet(t *testing.T) {
	arr := arrays.CreateNewArray("a", "b", "c", "d", "e")
	if result, _ := arr.Get(0); result != "a" {
		t.Errorf("Expected %v got %v", "a", result)
	}
	if result, _ := arr.Get(1); result != "b" {
		t.Errorf("Expected %v got %v", "b", result)
	}
	if result, _ := arr.Get(2); result != "c" {
		t.Errorf("Expected %v got %v", "c", result)
	}
	if result, _ := arr.Get(3); result != "d" {
		t.Errorf("Expected %v got %v", "d", result)
	}
	if result, _ := arr.Get(4); result != "e" {
		t.Errorf("Expected %v got %v", "e", result)
	}
	if _, err := arr.Get(5); err == nil {
		t.Errorf("Expected %v got %v", "array index out of bounds", err)
	}
}

func benchmarkGet(b *testing.B, arr *arrays.Array, testSize int) {

	for i := 0; i <= b.N; i++ {
		for j := 0; j < testSize; j++ {
			arr.Get(j)
		}
	}
}

func BenchmarkArrayGet100(b *testing.B) {
	b.StopTimer()
	arr := arrays.CreateNewArray()
	testSize := 100
	//adding 100 elements
	for i := 0; i < testSize; i++ {
		arr.Add(i)
	}
	b.StartTimer()
	benchmarkGet(b, arr, testSize)
}

func benchmarkRemoveAt(b *testing.B, arr *arrays.Array, testSize int) {

	for i := 0; i <= b.N; i++ {
		for j := testSize; j >= 0; j-- {
			arr.Get(j)
		}
	}
}

func BenchmarkArrayRemoveAt100(b *testing.B) {
	b.StopTimer()
	arr := arrays.CreateNewArray()
	testSize := 100
	//adding 100 elements
	for i := 0; i < testSize; i++ {
		arr.Add(i)
	}
	b.StartTimer()
	benchmarkRemoveAt(b, arr, testSize)
}
