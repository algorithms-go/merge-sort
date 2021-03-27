package main

import (
	"fmt"
	"github.com/algorithms-go/sorting/sort"
)

func main() {
	A := []int{2,4,5,7,1,2,3,6}
	fmt.Println(sort.MergeSort(A, 0, 3, 7))
}
