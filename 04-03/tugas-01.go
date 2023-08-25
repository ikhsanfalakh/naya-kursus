package main

import (
	"fmt"
)

/*
Buatlah sebuah program yang menentukan apakah kita dapat membuat sebuah segitiga dengan parameter yang di dapat (a, b,
c). Jika iya kembalikan nilai true dan false jika tidak.
*/

func isTriangle(a, b, c float64) bool {
	// Menggunakan ketentuan bahwa jumlah dua sisi apa pun harus lebih besar dari sisi ketiga
	return a+b > c && a+c > b && b+c > a
}

func main() {
	// Panjang sisi segitiga
	a := 3.0
	b := 5.0
	c := 7.0

	// Memeriksa apakah dapat membentuk segitiga
	result := isTriangle(a, b, c)
	if result {
		fmt.Println("Dapat membentuk segitiga")
	} else {
		fmt.Println("Tidak dapat membentuk segitiga")
	}
}
