package main

import "fmt"

/*
### Tugas
Buatlah program yang menterjemahkan angka romawi ke angka desimal.
```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

#### Example:
```
Input: s = "III"
Output: 3
Explanation: III = 3.

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
```
*/

var romawiNum = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romawiToNumber(romawi string) (result int, err string) {
	result = 0
	valuePrev := 0

	for i := len(romawi) - 1; i >= 0; i-- {
		value, isExist := romawiNum[string(romawi[i])]

		if isExist {
			if value < valuePrev {
				result -= value
			} else {
				result += value
			}
			valuePrev = value
		} else {
			result = 0
			err = string(romawi[i]) + " bukan angka romawi"
			break
		}
	}

	return result, err
}

func main() {
	var angkaromawi string
	fmt.Printf("Input Angka Romawi = ")
	fmt.Scan(&angkaromawi)

	fmt.Printf("Angka Decimal = ")
	value, err := romawiToNumber(angkaromawi)
	fmt.Println(value, err)
}
