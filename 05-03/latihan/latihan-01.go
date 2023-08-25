package latihan

import "sort"

func CariTengah(arr []int) int {
	if len(arr) < 3 {
		return -1
	}

	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	indexKe2 := sorted[1]

	for i, num := range arr {
		if num == indexKe2 {
			return i
		}
	}

	return -1
}
