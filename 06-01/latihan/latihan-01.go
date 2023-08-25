package latihan

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(inputStr string) []Tuple {
	charCount := make(map[rune]int)

	for _, char := range inputStr {
		charCount[char]++
	}

	result := make([]Tuple, len(charCount))
	i := 0
	for char, count := range charCount {
		result[i] = Tuple{Char: char, Count: count}
		i++
	}

	return result
}
