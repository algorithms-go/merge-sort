package main

import (
	"fmt"

	"github.com/algorithms-go/sorting/sort"
)

func main() {
	A := []int{5, 2, 4, 7, 1, 3, 2, 6}
	fmt.Println(sort.MergeSort(A))
	fmt.Println(sort.MergeSortFunc(A, func(a, b int) bool {
		return a < b
	}))
	fmt.Println(sort.InsertionMergeSortFunc(A, func(a, b int) bool {
		return a > b
	}))
	// explain merge in diagrams & animations
	// explain recursive call in diagrams & animations
}
