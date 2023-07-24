package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var angka1, angka2, angka3 int = 4, 8, 7
	var total int
	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(angka3)
	total = angka1 + angka2 + angka3
	fmt.Printf("total: %d\n", total)

	var angkadec1, angkadec2, tot_decimal float32
	angkadec1 = 5.8
	angkadec2 = 3.2
	tot_decimal = angkadec1 + angkadec2
	fmt.Println(tot_decimal)
	fmt.Printf("bilangan desimal1: %f\n", angkadec1)
	fmt.Printf("bilangan desimal2: %.2f\n", angkadec2)

	var str_var1, str_var2, str_var3 string = "Aku", "Belajar", "Golang"
	fmt.Println(str_var1 + str_var2 + str_var3)
	fmt.Printf("kalimat : %s %s %s\n", str_var1, str_var2, str_var3)
}
