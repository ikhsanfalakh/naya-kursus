package main

import "fmt"

/*
### Tugas
Buatlah program yang menerima input berupa array bertipe integer dan konversikanlah array tersebut menjadi bilangan desimal kemudian jumlahkan dengan 1 lalu kembalikan hasilnya berupa array bertipe integer juga.

#### Example:
~~~
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
~~~
*/
func incrementArray(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	newDigits := make([]int, n+1)
	newDigits[0] = 1

	return newDigits
}

func decrementArray(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] != 0 {
			digits[i]--
			if digits[0] == 0 {
				return digits[1:]
			}
			return digits
		}

		digits[i] = 9
	}

	return digits
}

func main() {
	digits1 := []int{2, 9}
	fmt.Println(digits1)
	output1 := incrementArray(digits1)
	fmt.Println(output1)

	digits2 := []int{9}
	fmt.Println(digits2)
	output2 := incrementArray(digits2)
	fmt.Println(output2)

	digits3 := []int{1, 0}
	fmt.Println(digits3)
	output3 := decrementArray(digits3)
	fmt.Println(output3)
}
