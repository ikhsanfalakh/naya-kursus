package main

import (
	"fmt"
	"strconv"
)

/*
### Latihan-1
Buatlah program yang menjumlahkan dua buah binary (a, b) dan kembalikan hasilnya.

#### Example:
```
Input: a = "11", b = "1"
Output: "100"

Input: a = "1010", b = "1011"
Output: "10101"
```
*/

func decToStrBinary(angka int64) string {
	result := strconv.FormatInt(angka, 2)

	return result
}

func main() {
	a := "111"
	b := "1"

	c, _ := strconv.ParseInt(a, 2, 32)
	d, _ := strconv.ParseInt(b, 2, 32)
	fmt.Println(c, d)
	total := c + d

	fmt.Println(decToStrBinary(total))

	e := "1010"
	f := "1011"

	g, _ := strconv.ParseInt(e, 2, 64)
	h, _ := strconv.ParseInt(f, 2, 64)
	fmt.Println(g, h)
	total2 := g + h

	fmt.Println(decToStrBinary(total2))
}
