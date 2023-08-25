package main

import "fmt"

/*
Buatlah program yang menentukan banyaknya angka dimulai dari start sampai dengan end (termasuk). Jika angka ke n mengandung angka 5, maka tidak di hitung.

1,9 -> 1,2,3,4,6,7,8,9 -> Result 8
4,17 -> 4,6,7,8,9,10,11,12,13,14,16,17 -> Result 12
*/

type Range struct {
	Start int
	End   int
}

func cekAdaLima(n int) bool {
	if n < 0 {
		n *= -1
	}
	if n%10 == 5 {
		return true
	}
	return false
}

func jmlAngkaTanpaLima(r Range) (count int, listAngka []int) {
	for i := r.Start; i <= r.End; i++ {
		if !cekAdaLima(i) {
			count++
			listAngka = append(listAngka, i)
		}
	}
	return count, listAngka
}

func main() {
	range1 := Range{Start: -35, End: 16}
	result1, listAngka1 := jmlAngkaTanpaLima(range1)
	fmt.Printf("Range: %d - %d\n", range1.Start, range1.End)
	fmt.Println("Angka:", listAngka1)
	fmt.Printf("Result: %d\n", result1)

	range2 := Range{Start: 4, End: 17}
	result2, listAngka2 := jmlAngkaTanpaLima(range2)
	fmt.Printf("Range: %d - %d\n", range2.Start, range2.End)
	fmt.Println("Angka:", listAngka2)
	fmt.Printf("Result: %d\n", result2)
}
