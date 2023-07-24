package main

import "fmt"

/*
### Latihan-2
Buatlah program yang menerima input nama, lalu tampilkan sapaan dengan
menggunakan nama tersebut!

Hasil:
```
Nama kamu siapa? <<nama>>
Selamat datang, <<nama>>.
```
*/
func main() {
	var name string

	fmt.Printf("Nama Kamu siapa? ")
	fmt.Scan(&name)
	fmt.Printf("Selamat datang, %s\n", name)
}
