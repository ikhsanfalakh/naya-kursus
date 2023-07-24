package main

import "fmt"

func main() {
	//operator aritmatika
	fmt.Println("==== penjumlahan ====")
	var a, b int = 4, 5
	var c int
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c = a + b
	fmt.Println("c = a + b\nc =", c)

	fmt.Println("==== pengurangan ====")
	c = b + a
	fmt.Println("c = b - a\nc =", c)

	fmt.Println("==== perkalian ====")
	c = a * b
	fmt.Println("c = a * b\nc =", c)

	fmt.Println("==== pembagian ====")
	var d int = 30
	c = d / b
	fmt.Println("d =", d)
	fmt.Println("c = d / b\nc =", c)

	fmt.Println("==== modulus ====")
	c = d % a
	fmt.Println("c = d % c\nc =", c)

	//operator assignment
	fmt.Println("==== increment ====")
	fmt.Println("a =", a)
	a++
	fmt.Println("a =", a)

	fmt.Println("==== decrement ====")
	fmt.Println("a =", a)
	a--
	fmt.Println("a =", a)

	fmt.Println("==== tambah sama dengan ====")
	fmt.Println("a =", a)
	a += 2
	fmt.Println("a += 2\na =", a)

	fmt.Println("==== kurang sama dengan ====")
	fmt.Println("a =", a)
	a -= 2
	fmt.Println("a -= 2\na =", a)

	fmt.Println("==== kali sama dengan ====")
	fmt.Println("a =", a)
	a *= 2
	fmt.Println("a *= 2\na =", a)

	fmt.Println("==== bagi sama dengan ====")
	fmt.Println("a =", a)
	a /= 2
	fmt.Println("a /= 2\na =", a)

	fmt.Println("==== modulo sama dengan ====")
	fmt.Println("a =", a)
	a %= 2
	fmt.Println("a %= 2\na =", a)

	//operator perbandingan
	fmt.Println("==== operator perbandingan ====")
	var e bool
	fmt.Println("e =", e)
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	fmt.Println("==== equal ====")
	e = a == b
	fmt.Println("a == b -->", e)

	fmt.Println("==== tidak samadengan ====")
	e = a != b
	fmt.Println("a != b -->", e)

	fmt.Println("==== kurang dari ====")
	e = a < b
	fmt.Println("a < b -->", e)

	fmt.Println("==== kurang dari samadengan ====")
	e = a <= b
	fmt.Println("a <= b -->", e)

	fmt.Println("==== lebih dari ====")
	e = a > b
	fmt.Println("a > b -->", e)

	fmt.Println("==== lebih dari samadengan ====")
	e = a > b
	fmt.Println("a > b -->", e)

	//operator logika
	fmt.Println("==== logika TRUE AND TRUE ====")
	e = true && true
	fmt.Println("true && true =", e)

	fmt.Println("==== logika TRUE AND FALSE ====")
	e = true && false
	fmt.Println("true && false =", e)

	fmt.Println("==== logika TRUE OR TRUE ====")
	e = true || true
	fmt.Println("true || true =", e)

	fmt.Println("==== logika TRUE OR FALSE ====")
	e = true || false
	fmt.Println("true || false =", e)

	fmt.Println("==== logika FALSE OR FALSE ====")
	e = false || false
	fmt.Println("false || false =", e)

	fmt.Println("==== logika NEGASI TRUE ====")
	e = !true
	fmt.Println("!true =", e)

	fmt.Println("==== logika NEGASI FALSE ====")
	e = !false
	fmt.Println("!false =", e)

}
