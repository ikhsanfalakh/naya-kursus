package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)
	fmt.Println("")

	numberA = 5
	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)
	fmt.Println("")

	var numberC int = *numberB
	fmt.Println("numberC (value)   :", numberC)
	fmt.Println("numberC (address) :", &numberC)
	fmt.Println("")

	var number = 4
	fmt.Println("before (value)   :", number)
	fmt.Println("before (address) :", &number)
	change(&number, 10)
	fmt.Println("after (value)    :", number)
	fmt.Println("after (address)  :", &number)
}

func change(original *int, value int) {
	*original = value
	fmt.Println("(address) : ", &original)
}
