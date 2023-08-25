package main

import (
	"fmt"
)

type Antrian []int

func (q *Antrian) in_antrian(item int) {
	*q = append(*q, item)
}

func processAntrian(antrian []int) []int {
	resultAntrian := Antrian{}

	for _, num := range antrian {
		if num%2 == 0 {
			resultAntrian.in_antrian(num / 2)
		} else {
			resultAntrian.in_antrian(num * 2)
		}
	}

	result := []int(resultAntrian)
	return result
}

func main() {
	inputAntrian1 := []int{1, 2, 3, 4, 5}
	outputAntrian1 := processAntrian(inputAntrian1)
	fmt.Println("Output 1:", outputAntrian1) // Output 1: [2 1 6 2 10]

	inputAntrian2 := []int{10, 20, 30, 40, 50}
	outputAntrian2 := processAntrian(inputAntrian2)
	fmt.Println("Output 2:", outputAntrian2) // Output 2: [5 10 15 20 25]
}
