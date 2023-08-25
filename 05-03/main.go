package main

import (
	"0503/latihan"
	"0503/tugas"
	"fmt"

	_ "github.com/ngamux/ngamux"
)

func main() {
	fmt.Println("===Latihan===")
	array1 := []int{2, 3, 1}
	middleIndex1 := latihan.CariTengah(array1)
	fmt.Printf("%v nilai tengah adalah index ke %d\n", array1, middleIndex1)

	array2 := []int{24, 12, 18}
	middleIndex2 := latihan.CariTengah(array2)
	fmt.Printf("%v nilai tengah adalah index ke %d\n", array2, middleIndex2)

	fmt.Println("===Tugas===")
	umur1 := []int{1, 2, 10, 8}
	result1 := tugas.CariUmur2Tertua(umur1)
	fmt.Printf("%v 2 yang tertua %v\n", umur1, result1)

	umur2 := []int{1, 5, 87, 45, 8, 8}
	result2 := tugas.CariUmur2Tertua(umur2)
	fmt.Printf("%v 2 yang tertua %v\n", umur2, result2)

	umur3 := []int{1, 3, 10, 0}
	result3 := tugas.CariUmur2Tertua(umur3)
	fmt.Printf("%v 2 yang tertua %v\n", umur3, result3)

}
