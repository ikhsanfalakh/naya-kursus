package main

import (
	"fmt"
	"strconv"
)

/*
#Latihan-3
Buatlah program aplikasi kalkulator sederhana yang meminta user input
Inputan untuk  memasukan jenis operator aritmatika yang dipilih
Inputan untuk memasukan nilai dari variabel pertama
Inputan untuk memasukan nilai dari variabel kedua
Menampilkan hasil dari operasi variabel pertama dan kedua
*/

func main() {
	var operator, jns_angka, input1, input2 string
	var angka_int1, angka_int2, tot_int int
	var angka_dec1, angka_dec2, tot_dec float64
	var err_ error

	//input operator
	for {
		fmt.Printf("Masukkan operator [ + | - | * | / ] = ")
		_, err := fmt.Scan(&operator)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		if (operator == "+") ||
			(operator == "-") ||
			(operator == "*") ||
			(operator == "/") {
			break
		} else {
			fmt.Println("Bukan Operator!")
			continue
		}
		break
	}

	//input jenis angka
	for {
		fmt.Printf("Masukkan Jenis Number [integer | decimal] = ")
		_, err := fmt.Scan(&jns_angka)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		if (jns_angka == "integer") ||
			(jns_angka == "decimal") {
			break
		} else {
			fmt.Println("Jenis Number tidak ada!")
			continue
		}
		break
	}

	//input angka pertama
	for {
		fmt.Printf("Masukkan Angka Pertama = ")
		_, err := fmt.Scan(&input1)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}

		if jns_angka == "integer" {
			angka_int1, err = strconv.Atoi(input1)
			if err != nil {
				fmt.Println(input1, " Bukan Angka!")
				continue
			}
		} else if jns_angka == "decimal" {
			angka_dec1, err_ = strconv.ParseFloat(input1, 64)
			if err_ != nil {
				fmt.Println(input1, " Bukan Angka!")
				continue
			}
		}
		break
	}

	//input angka kedua
	for {
		fmt.Printf("Masukkan Angka Kedua = ")
		_, err := fmt.Scan(&input2)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}

		if jns_angka == "integer" {
			angka_int2, err = strconv.Atoi(input2)
			if err != nil {
				fmt.Println(input2, " Bukan Angka!")
				continue
			}
		} else if jns_angka == "decimal" {
			angka_dec2, err = strconv.ParseFloat(input2, 64)
			if err != nil {
				fmt.Println(input2, " Bukan Angka!")
				continue
			}
		}
		break
	}

	//proses hitung
	if jns_angka == "integer" {
		if operator == "+" {
			tot_int = angka_int1 + angka_int2
		} else if operator == "-" {
			tot_int = angka_int1 - angka_int2
		} else if operator == "*" {
			tot_int = angka_int1 * angka_int2
		} else if operator == "/" {
			tot_int = angka_int1 / angka_int2
		}

		fmt.Println("Angka Pertama =", angka_int1)
		fmt.Println("Angka Kedua =", angka_int2)
		fmt.Println("Total =", angka_int1, operator, angka_int2)
		fmt.Println("Total =", tot_int)
	} else if jns_angka == "decimal" {
		if operator == "+" {
			tot_dec = angka_dec1 + angka_dec2
		} else if operator == "-" {
			tot_dec = angka_dec1 - angka_dec2
		} else if operator == "*" {
			tot_dec = angka_dec1 * angka_dec2
		} else if operator == "/" {
			tot_dec = angka_dec1 / angka_dec2
		}

		fmt.Println("Angka Pertama =", angka_dec1)
		fmt.Println("Angka Kedua =", angka_dec2)
		fmt.Printf("Total = %.2f %s %.2f\n", angka_dec1, operator, angka_dec2)
		fmt.Printf("Total = %.2f\n", tot_dec)
	}

}
