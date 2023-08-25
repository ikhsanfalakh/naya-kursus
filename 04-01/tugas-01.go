package main

import (
	"fmt"
	"math"
)

/*
### Tugas
Buatlah program yang mengecek apakah inputan n bertipe int merupakan hasil dari pangkat 2 atau bukan. Jika iya kembalikan
true dan false jika bukan.

#### Example:
```
Input: n = 1
Output: true

Input: n = 3
Output: false
```
*/

func isPowerOfTwo(angka int, hasil *bool) {
	num := math.Sqrt(float64(angka))
	fmt.Println(num)
	dpn, dec := math.Modf(num)
	fmt.Println(dpn)
	fmt.Println(dec)
	*hasil = (dec == 0)
}

func main() {
	var num int
	var angkaPangkat bool

	fmt.Printf("Inputkan Angka = ")
	fmt.Scan(&num)

	isPowerOfTwo(num, &angkaPangkat)

	if angkaPangkat {
		fmt.Println(num, "adalah pangkat dari 2")
	} else {
		fmt.Println(num, "bukan pangkat dari 2")
	}
}
