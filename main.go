package main

import (
	"fmt"
	"github.com/algorithms-go/sorting/sort"
)

func main() {
	//A := []int{2, 4, 5, 7, 1, 2, 3, 6}
	//A := []int{5, 2, 4, 7, 1, 3, 2, 6}
	A := []int{8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(sort.MergeSort(A))
	// call merge on multiple sets to prove it works
	// explain merge in diagrams & animations
	// explain recursive call in diagrams & animations
}
