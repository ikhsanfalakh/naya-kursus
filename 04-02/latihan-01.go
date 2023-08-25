package main

import "fmt"

/*
### Latihan-1

Buatlah program yang secara berulang menambah semua digitnya sampai hanya tersisa 1 digit saja dan return lah nilai dari
digit tersebut.

~~~
Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
~~~
*/

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
		fmt.Println(num)
	}
	return num
}

func main() {
	num := 3867
	result := addDigits(num)
	fmt.Println(num, result)
}
