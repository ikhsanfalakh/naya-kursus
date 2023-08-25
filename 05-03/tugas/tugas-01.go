package tugas

import "sort"

func CariUmur2Tertua(ages []int) []int {
	if len(ages) < 2 {
		return nil
	}

	sorted := make([]int, len(ages))
	copy(sorted, ages)
	sort.Ints(sorted)

	return []int{sorted[len(sorted)-2], sorted[len(sorted)-1]}
}
