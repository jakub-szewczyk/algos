package iterative_sorts

// Time complexity
// O(n²)
// Θ(n²)
// Ω(n)

// Space complexity
// O(1)

// Stable
// Destructive

// Builds the sorted array one element at a time, shifting larger elements as needed.
// Best for: small datasets, nearly sorted arrays (O(n) best case), and when minimal extra space is needed.
func InsertionSort(xs []int) []int {
	for i := 1; i < len(xs); i++ {
		for j := i; j > 0 && xs[j] < xs[j-1]; j-- {
			xs[j-1], xs[j] = xs[j], xs[j-1]
		}
	}
	return xs
}
