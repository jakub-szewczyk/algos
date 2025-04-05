package practice

func Factorial(x int) int {
	if x < 2 {
		return 1
	}
	return x * Factorial(x-1)
}
