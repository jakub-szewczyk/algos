package search

// Time complexity
// O(log n)
// Θ(log n)
// Ω(1)

// Space complexity
// O(1)
// Θ(1)
// Ω(1)

// Non-destructive

// An efficient searching algorithm that works on sorted arrays by repeatedly dividing the search space in half.
// Instead of checking each element one by one (O(n) like linear search), it cuts the search space in half at each step, achieving O(log n) time complexity.
func BinarySearch(v int, xs []int) (int, int) {
	l := 0
	h := len(xs) - 1

	for l <= h {
		i := (l + h) / 2

		if v < xs[i] {
			h = i - 1
		} else if v > xs[i] {
			l = i + 1
		} else {
			return i, v
		}
	}

	return -1, -1
}
