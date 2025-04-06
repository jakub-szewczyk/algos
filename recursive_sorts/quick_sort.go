package recursive_sorts

// Time complexity
// O(n²)
// Θ(n log n)
// Ω(n log n)

// Space complexity
// O(n)
// Θ(log n)
// Ω(log n)

// Not stable
// Destructive/non-destructive

// A divide-and-conquer algorithm that partitions the array around a pivot, sorting in place.
// Best for: large datasets, arrays, and general-purpose sorting with O(n log n) average time.
// Downside: worst-case O(n²) if the pivot is poorly chosen (avoided with randomization).
func QuickSort(xs []int) []int {
	if len(xs) < 2 {
		return xs
	}

	i := len(xs) - 1
	p := xs[i]

	left, right := []int{}, []int{}
	for _, x := range xs[:i] {
		if x < p {
			left = append(left, x)
		} else {
			right = append(right, x)
		}
	}

	ys, zs := QuickSort(left), QuickSort(right)

	final := []int{}
	final = append(final, ys...)
	final = append(final, p)
	final = append(final, zs...)

	return final
}
