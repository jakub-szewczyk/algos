package recursive_sorts

func merge(xs []int, ys []int) []int {
	zs := []int{}

	for len(xs) > 0 && len(ys) > 0 {
		if xs[0] < ys[0] {
			zs = append(zs, xs[0])
			xs = xs[1:]
		} else {
			zs = append(zs, ys[0])
			ys = ys[1:]
		}
	}

	zs = append(zs, xs...)
	zs = append(zs, ys...)

	return zs
}

// Time complexity
// O(n log n)
// Θ(n log n)
// Ω(n log n)

// Space complexity
// O(n)

// Stable
// Non-destructive

// A divide-and-conquer algorithm with O(n log n) time complexity, splitting the array recursively and merging sorted halves.
// JavaScript engines use it under the hood when `Array.prototype.sort()` method is invoked.
// Great for: linked lists, stable sorting, parallel processing, and guaranteed performance.
// Downside: requires extra O(n) space.
func MergeSort(xs []int) []int {
	if len(xs) < 2 {
		return xs
	}

	i := len(xs) / 2

	ys, zs := MergeSort(xs[:i]), MergeSort(xs[i:])

	return merge(ys, zs)
}
