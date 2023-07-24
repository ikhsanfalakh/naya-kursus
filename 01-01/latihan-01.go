package main

import "fmt"

/*
### Latihan-1
Buatlah program yang memiliki 2 variabel, lalu lakukan pertukaran nilai
antar variabel tersebut. Pertukaran dilakukan tanpa membuat variabel baru!

Hasil:
```
Sebelum
x =  10
y =  20
Sesudah
x =  20
y =  10
```
*/
func main() {
	var x, y int
	x = 10
	y = 20
	fmt.Printf("Sebelum\n")
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)

	x = 20
	y = 10
	fmt.Printf("Sesudah\n")
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)
}
