package main

import (
	"fmt"
)

/*
Buatlah program yang bertujuan untuk mengecek apakah karakter terakhir dari string a sama dengan string b. Jika iya kembalikan nilai true, dan jika tidak sama kembalikan nilai false.

solution("abc", "bc") // returns true
solution("abc", "d") // returns false
*/

func checkLastChar(a, b string) (result bool) {
	if len(a) < len(b) {
		return
	}

	fmt.Println(a[len(a)-len(b):])
	fmt.Println(len(a))
	fmt.Println(len(b))
	substringFromLast := a[len(a)-len(b):]
	return substringFromLast == b
}

func main() {
	var stringA, stringB string
	fmt.Printf("Masukkan Kata = ") //indonesia
	fmt.Scan(&stringA)
	fmt.Printf("Cek karakter = ") //ria
	fmt.Scan(&stringB)

	result := checkLastChar(stringA, stringB)

	if result {
		fmt.Println("Karakter terakhir sama")
	} else {
		fmt.Println("Karakter terakhir tidak sama")
	}
}
