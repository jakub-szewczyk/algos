package main

import (
	"fmt"

	"github.com/jakub-szewczyk/algos/iterative_sorts"
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
}
