package non_comparison_sorts

import (
	"math"
	"slices"
	"strconv"
)

func RadixSort(xs []int) []int {
	buckets := [10][]int{}

	for i := 0; i < len(strconv.Itoa(slices.Max(xs))); i++ {
		for _, x := range xs {
			digit := 0

			if i == 0 {
				digit = x % 10
			} else {
				digit = int(math.Floor(float64(x/(int(math.Pow10(i)))))) % 10
			}

			buckets[digit] = append(buckets[digit], x)
		}

		xs = nil

		for j, bucket := range buckets {
			xs = append(xs, bucket...)
			buckets[j] = nil
		}
	}

	return xs
}
