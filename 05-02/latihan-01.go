package main

import (
	"fmt"
)

/*
Buatlah program yang menentukan nilai maksimum dan minimun dari sebuah array.

[1,2,3,4,5] --> [1,5]
[2334454,5] --> [5,2334454]
[1]         --> [1,1]
*/

type MinMax struct {
	Min int
	Max int
}

func cariMinMax(arr []int) MinMax {
	if len(arr) == 0 {
		return MinMax{}
	}

	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return MinMax{Min: min, Max: max}
}

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	result1 := cariMinMax(array1)
	fmt.Printf("Array: %v, Min: %d, Max: %d\n", array1, result1.Min, result1.Max)

	array2 := []int{2334454, 5}
	result2 := cariMinMax(array2)
	fmt.Printf("Array: %v, Min: %d, Max: %d\n", array2, result2.Min, result2.Max)

	array3 := []int{1}
	result3 := cariMinMax(array3)
	fmt.Printf("Array: %v, Min: %d, Max: %d\n", array3, result3.Min, result3.Max)
}
