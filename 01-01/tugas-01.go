package main

import (
	"fmt"
	"strconv"
)

/*
### Tugas
Buatlah program yang menerima input user berupa angka, lalu tampilkan hasil dari faktorial angka tersebut. Selama yang di inputkan bukan angka (Not a Number) tampilkanlah pesan lalu ulangi proses input kembali.

Hasil:
```
Enter a number: number
Please enter a number!
Enter a number: this is a number
Please enter a number!
Enter a number: 5
120
```
*/
func main() {
	var input string
	var number, hasil int

	//check is a number
	for {
		fmt.Printf("Enter a number: ")
		_, err := fmt.Scanln(&input)
		fmt.Printf("Please enter a number!\n")
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}

		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a number: this is Not a number")
			continue
		}

		break
	}
	fmt.Println("Enter a number: this is a number")
	fmt.Printf("Enter a number: %d\n", number)
	//hitung number factorial
	if number == 0 {
		hasil = 1
	} else {
		hasil = 1
		for i := 1; i <= number; i++ {
			hasil *= i
		}
	}
	fmt.Printf("Result Factorial: %d\n", hasil)
}
