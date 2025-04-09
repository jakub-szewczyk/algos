package search

// Time complexity
// O(n)
// Θ(n)
// Ω(1)

// Space complexity
// O(1)
// Θ(1)
// Ω(1)

// Non-destructive

// A straightforward searching algorithm that checks each element one by one until it finds the target or reaches the end of the array.
// Unlike binary search, it doesn’t require sorted data but is much slower for large datasets.
func LinearSearch(v int, xs []int) (int, int) {
	for i, x := range xs {
		if v == x {
			return i, v
		}
	}
	return -1, -1
}
