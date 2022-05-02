//This module implements the Dynamic Array Concept of DS althought it is much more efficeient to use slices
package arrays

import (
	"errors"
	"log"
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
		elements: make([]interface{}, 0),
		length:   -1, //inital index of array with no elements
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
	if (arr.capacity - (arr.length + 1)) <= len(values) { // arr.length represents the last "index"
		//increase capacity
		arr.growBy(len(values))
		newElement := make([]interface{}, arr.capacity)
		//copy old array elemets to new array; old length remains same
		for i := 0; i <= arr.length; i++ {
			newElement[i] = arr.elements[i]
		}
		arr.elements = newElement
	}
	for _, val := range values {
		arr.length++
		arr.elements[arr.length] = val
	}
}

//RemoveAt removes the array at a specific index
func (arr *Array) RemoveAt(index int) (interface{}, error) {
	if arr.IsEmpty() || index > arr.length {
		return nil, errors.New("array Index out of bounds")
	}
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
	log.Println(arr.elements, arr.capacity, arr.length)
	//shrink capacity
	arr.shrinkBy()
	return data, nil
}

func (arr *Array) Remove(element interface{}) error {

	if arr.Contains(element) {
		newElementArray := make([]interface{}, arr.capacity)
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
		return nil
	}
	return errors.New("element not found")
}

//IndexOf retuns index of element otherwise returns -1 and error message
func (arr *Array) IndexOf(element interface{}) (int, error) {
	for i := 0; i <= arr.length; i++ {
		if arr.elements[i] == element {
			return i, nil
		}
	}
	return -1, errors.New("element not found")
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
func (arr *Array) IsEmpty() bool { return arr.length == -1 }

//Get returns an element from index number
func (arr *Array) Get(index int) (interface{}, error) {
	if index > arr.length {
		return nil, errors.New("array index out of bounds")
	}
	return arr.elements[index], nil
}

//Set assigns a the value to the index passed in the parameters
func (arr *Array) Set(index int, val interface{}) error {
	if index > arr.length {
		return errors.New("array index out of bounds")
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
	arr.capacity = int(growthFactor) * (arr.length + valueLength)
}

func (arr *Array) shrinkBy() {
	if arr.length == 0 {
		arr.capacity = 2
	} else if arr.capacity/arr.length > 3.0 {
		arr.capacity = arr.capacity - int(float32(arr.capacity)*shrinkFactor)
	} else {
		return
	}
	newElementArray := make([]interface{}, arr.capacity)
	copy(newElementArray, arr.elements)
	arr.elements = newElementArray
}
