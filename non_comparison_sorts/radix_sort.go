package non_comparison_sorts

import (
	"math"
	"slices"
	"strconv"
)

// Time complexity
// O(nk)
// Θ(nk)
// Ω(nk)

// Space complexity
// O(n + k)
// Θ(n + k)
// Ω(n + k)

// Stable
// Destructive/non-destructive

// A non-comparative sorting algorithm that sorts numbers digit by digit (least significant to most significant) using a stable sub-sorting method.
// It works best on integers and fixed-length strings.
// Best for:
// - sorting large numbers efficiently (O(nk) time complexity, where k is the number of digits),
// - fixed-length strings or data with a known range.
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
