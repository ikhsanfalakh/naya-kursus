package main

import (
	"fmt"
	"math"
)

/*
Buatlah program untuk menghitung:
- keliling dan luas bangun datar (minimum 2)
- luas, keliling dan volume bangun ruang (minimum 2)
*/

func main() {
	//Menu

	//luas dan keliling persegi panjang
	var panjang, lebar, luas_persegipanjang, keliling_persegipanjang int
	fmt.Println("Hitung Luas dan Keliling Persegi Panjang")
	//input panjang
	for {
		fmt.Printf("Panjang = ")
		_, err := fmt.Scan(&panjang)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}
	//input lebar
	for {
		fmt.Printf("Lebar = ")
		_, err := fmt.Scan(&lebar)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}

	//hitung luas persegi panjang
	luas_persegipanjang = panjang * lebar
	fmt.Println("Luas =", luas_persegipanjang)

	//hitung keliling persegi panjang
	keliling_persegipanjang = 2 * (panjang + lebar)
	fmt.Println("Keliling =", keliling_persegipanjang)

	//luas dan keliling segitiga sama sisi
	var alas, tinggi, keliling_segitigasamasisi int
	var luas_segitigasamasisi float32
	fmt.Println("Hitung Luas dan Keliling Segitiga Sama Sisi")
	//input alas
	for {
		fmt.Printf("Alas = ")
		_, err := fmt.Scan(&alas)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}
	//input tinggi
	for {
		fmt.Printf("Tinggi = ")
		_, err := fmt.Scan(&tinggi)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}
	//hitung luas Segitiga Sama Sisi
	luas_segitigasamasisi = float32(alas) * float32(tinggi) / 2
	fmt.Println("Luas =", luas_segitigasamasisi)

	//hitung keliling Segitiga Sama Sisi
	keliling_segitigasamasisi = 3 * alas
	fmt.Println("Keliling =", keliling_segitigasamasisi)

	//===================================================
	//luas, keliling dan volume Balok
	var luas_balok, keliling_balok, volume_balok int
	fmt.Println("Hitung Luas dan Volume Balok")
	//input panjang
	for {
		fmt.Printf("Panjang = ")
		_, err := fmt.Scan(&panjang)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}
	//input lebar
	for {
		fmt.Printf("Lebar = ")
		_, err := fmt.Scan(&lebar)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}
	//input tinggi
	for {
		fmt.Printf("Tinggi = ")
		_, err := fmt.Scan(&tinggi)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}

	//hitung luas balok
	luas_balok = 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
	fmt.Println("Luas Balok =", luas_balok)

	//hitung keliling balok
	keliling_balok = 4 * (panjang + lebar + tinggi)
	fmt.Println("Keliling Balok =", keliling_balok)

	//hitung volume balok
	volume_balok = panjang * lebar * tinggi
	fmt.Println("Volume Balok =", volume_balok)

	//luas, keliling dan volume Tabung
	var jarijari int
	var luas_tabung, keliling_tabung, volume_tabung float64
	fmt.Println("Hitung Luas dan Volume Balok")
	//input jari-jari
	for {
		fmt.Printf("Jari-Jari = ")
		_, err := fmt.Scan(&jarijari)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}
	//input tinggi
	for {
		fmt.Printf("Tinggi = ")
		_, err := fmt.Scan(&tinggi)
		if err != nil {
			fmt.Println("Tidak Boleh Kosong")
			continue
		}
		break
	}

	//hitung luas tabung
	luas_tabung = 2 * math.Phi * math.Sqrt(float64(jarijari))
	fmt.Printf("Luas Tabung = %.2f\n", luas_tabung)

	//hitung keliling alas tabung
	keliling_tabung = 2 * math.Phi * float64(jarijari)
	fmt.Printf("Keliling Alas Tabung = %.2f\n", keliling_tabung)

	//hitung volume tabung
	volume_tabung = math.Phi * math.Sqrt(float64(jarijari)) * float64(tinggi)
	fmt.Printf("Volume Tabung = %.2f\n", volume_tabung)
}
