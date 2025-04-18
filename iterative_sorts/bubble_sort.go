package iterative_sorts

// Time complexity
// O(n²)
// Θ(n²)
// Ω(n)

// Space complexity
// O(1)

// Stable (maintains the relative order of equal elements)
// Destructive (mutates the input)

// A simple but inefficient O(n²) sorting algorithm that repeatedly swaps adjacent elements until sorted.
// Used for: teaching purposes, detecting nearly sorted data, and small datasets.
// Not practical for large inputs.
func BubbleSort(xs []int) []int {
	for i := 0; i < len(xs); i++ {
		swapped := false

		for j := 1; j < len(xs)-i; j++ {
			if xs[j-1] > xs[j] {
				xs[j-1], xs[j] = xs[j], xs[j-1]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return xs
}
