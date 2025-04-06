package recursive_sorts

func merge(xs []int, ys []int) []int {
	zs := []int{}
	length := len(xs) + len(ys)

	for len(zs) < length {
		if len(xs) > 0 && len(ys) == 0 {
			zs = append(zs, xs...)
			break
		}
		if len(xs) == 0 && len(ys) > 0 {
			zs = append(zs, ys...)
			break
		}

		if xs[0] < ys[0] {
			zs = append(zs, xs[0])
			xs = xs[1:]
		} else {
			zs = append(zs, ys[0])
			ys = ys[1:]
		}
	}

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
func MergeSort(xs []int) []int {
	if len(xs) < 2 {
		return xs
	}

	i := len(xs) / 2

	ys := MergeSort(xs[:i])
	zs := MergeSort(xs[i:])

	return merge(ys, zs)
}
