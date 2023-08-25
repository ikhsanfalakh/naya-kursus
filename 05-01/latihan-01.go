package main

import (
	"fmt"
)

type StringPair struct {
	First  string
	Second string
}

func (sp StringPair) CheckLastCharacter() bool {
	if len(sp.First) == 0 || len(sp.Second) == 0 {
		return false
	}

	lastCharFirst := sp.First[len(sp.First)-1]
	lastCharSecond := sp.Second[len(sp.Second)-1]

	return lastCharFirst == lastCharSecond
}

func main() {
	str1 := StringPair{First: "hello", Second: "world"}
	str2 := StringPair{First: "apple", Second: "banana"}

	result1 := str1.CheckLastCharacter()
	if result1 {
		fmt.Println("Karakter terakhir sama untuk str1")
	} else {
		fmt.Println("Karakter terakhir tidak sama untuk str1")
	}

	result2 := str2.CheckLastCharacter()
	if result2 {
		fmt.Println("Karakter terakhir sama untuk str2")
	} else {
		fmt.Println("Karakter terakhir tidak sama untuk str2")
	}
}
