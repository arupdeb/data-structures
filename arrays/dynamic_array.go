//This module implements the Dynamic Array Concept of DS althought it is much more efficeient to use slices
package arrays

import (
	"errors"
)

//Array is a struct of slice representing a dynamic Array and will not use the append or any inbuilt function
type Array struct {
	elements []interface{} //inital value static array
	length   int           // length user thinks array is; also represents the last index
	capacity int           // Actual Array size
}

//growthFactor & shrinkFactor determines the change in capacity of the array
const (
	growthFactor = float32(2)
	shrinkFactor = float32(0.25)
)

func CreateNewArray(values ...interface{}) *Array {
	//initalize the new array
	arr := &Array{
		elements: make([]interface{}, 0, len(values)),
		length:   0,
		capacity: len(values),
	}
	if len(values) != 0 {
		//add fucntion call
		arr.Add(values...)
	}
	return arr
}

//Add takes a number of values to add to the dynamic Array
func (arr *Array) Add(values ...interface{}) {
	if (arr.capacity - arr.length + 1) <= len(values) { // arr.length represents the last "index"
		//increase capacity
		arr.growBy(len(values))
		newElement := make([]interface{}, arr.capacity, arr.capacity)
		//copy old array elemets to new array; old length remains same
		for i := 0; i <= arr.length; i++ {
			newElement[i] = arr.elements[i]
		}
	}
	for _, val := range values {
		arr.elements[arr.length] = val
		arr.length++
	}
}

func (arr *Array) RemoveAt(index int) interface{} {
	data := arr.elements[index]
	if index != arr.length {
		for i := 0; i < arr.length; i++ {
			if i >= index {
				arr.elements[i] = arr.elements[i+1]
			}
		}
	}
	arr.elements[arr.length] = nil
	arr.length--
	//shrink capacity
	if arr.capacity/arr.length > 3.0 {
		arr.shrinkBy()
	}
	return data
}

func (arr *Array) Remove(element interface{}) {

	newElementArray := make([]interface{}, arr.capacity, arr.capacity)
	for i, j := 0, 0; i <= arr.length; i, j = i+1, j+1 {
		if arr.elements[i] == element {
			j--
		} else {
			newElementArray[j] = arr.elements[i]
		}
	}
	arr.elements = newElementArray
	arr.length--
	//check for shirnk factor
	if arr.capacity/arr.length > 3.0 {
		arr.shrinkBy()
	}

}

//IndexOf retuns index of element otherwise returns -1 and error message
func (arr *Array) IndexOf(element interface{}) (int, error) {
	for i := 0; i <= arr.length; i++ {
		if arr.elements[i] == element {
			return i, nil
		}
	}
	return -1, errors.New("Element not found")
}

//Contains retuns true if element found otherwise false
func (arr *Array) Contains(element interface{}) bool {
	for i := 0; i <= arr.length; i++ {
		if arr.elements[i] == element {
			return true
		}
	}
	return false
}

//Size returns actual lenght of the array (lastindex +1)
func (arr *Array) Size() int { return arr.length + 1 }

//IsEmpty return true if Array is empty otherwise returns false
func (arr *Array) IsEmpty() bool { return arr.length == 0 }

//Get returns an element from index number
func (arr *Array) Get(index int) (interface{}, error) {
	if index > arr.length {
		return nil, errors.New("Array index out of bounds")
	}
	return arr.elements[index], nil
}

//Set assigns a the value to the index passed in the parameters
func (arr *Array) Set(index int, val interface{}) error {
	if index > arr.length {
		return errors.New("Array index out of bounds")
	}
	arr.elements[index] = val
	return nil
}

//Clear removes all the elements from the array
func (arr *Array) Clear() {
	for i := 0; i <= arr.length; i++ {
		arr.elements[i] = nil
	}
	arr.length = 0
}

//growBy increases the capacity based on the value to be added
func (arr *Array) growBy(valueLength int) {
	arr.capacity = int(growthFactor) * (arr.capacity + valueLength)
}

func (arr *Array) shrinkBy() {
	arr.capacity = arr.capacity - int(float32(arr.capacity)*shrinkFactor)
}
