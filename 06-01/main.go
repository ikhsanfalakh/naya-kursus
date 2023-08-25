package main

import (
	"0601/latihan"
	"fmt"
)

/*
Buatlah program yang menjumlah kemunculan dari setiap karakter pada string yang diberikan.
OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

	type Tuple struct {
	    Char  rune
	    Count int
	}
*/

func main() {
	kalimat1 := "abracadabra"
	tuple1 := latihan.OrderedCount(kalimat1)
	fmt.Println("kalimat =", kalimat1)
	for _, tuple := range tuple1 {
		fmt.Printf("Tuple{'%c', %d}\n", tuple.Char, tuple.Count)
	}
	fmt.Println()

	kalimat2 := "ikhsanfalakh@Ikhsans-MacBook-Pro"
	tuple2 := latihan.OrderedCount(kalimat2)
	fmt.Println("kalimat =", kalimat2)
	for _, tuple := range tuple2 {
		fmt.Printf("Tuple{'%c', %d}\n", tuple.Char, tuple.Count)
	}

	//var a rune = 112
	fmt.Printf("%c", 112)
}
