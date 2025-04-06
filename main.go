package main

import (
	"fmt"

	"github.com/jakub-szewczyk/algos/iterative_sorts"
	"github.com/jakub-szewczyk/algos/recursive_sorts"
)

func main() {
	fmt.Println("ITERATIVE SORTS")
	fmt.Println("----------------------------------------------------------------------------------------------------")

	xs := []int{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	fmt.Println("Input", xs)
	fmt.Println()
	fmt.Println("Bubble sort", iterative_sorts.BubbleSort(xs))
	fmt.Println()

	xs = []int{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	fmt.Println("Input", xs)
	fmt.Println()
	fmt.Println("Insertion sort", iterative_sorts.InsertionSort(xs))
	fmt.Println()

	fmt.Println("RECURSIVE SORTS")
	fmt.Println("----------------------------------------------------------------------------------------------------")

	// ys := []interface{}{1, []interface{}{2, 3}, 4, []interface{}{5, []interface{}{6, 7, []interface{}{1, 1}}}}
	// fmt.Println("Input", ys)
	// fmt.Println()
	// fmt.Println("Nested addition", practice.NestedAdd(ys))
	// fmt.Println()

	// x := 10
	// fmt.Println("Input", x)
	// fmt.Println()
	// fmt.Println("Factorial", practice.Factorial(10))
	// fmt.Println()

	xs = []int{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	fmt.Println("Input", xs)
	fmt.Println()
	fmt.Println("Merge sort", recursive_sorts.MergeSort(xs))
	fmt.Println()

	xs = []int{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	fmt.Println("Input", xs)
	fmt.Println()
	fmt.Println("Quick sort", recursive_sorts.QuickSort(xs))
	fmt.Println()
}
