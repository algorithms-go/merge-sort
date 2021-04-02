package main

import (
	"fmt"
	"math/rand"
	"time"

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

	benchmark(100_000)
}

func benchmark(n int) {
	var A []int
	for i := 0; i < n; i++ {
		A = append(A, rand.Intn(n))
	}
	start := time.Now()
	sort.InsertionSortFunc(A, func(a, b int) bool {
		return a < b
	})
	fmt.Println("elapsed for insertion sort:", time.Now().Sub(start))

	start = time.Now()
	sort.MergeSortFunc(A, func(a, b int) bool {
		return a < b
	})
	fmt.Println("elapsed for merge sort:", time.Now().Sub(start))

	start = time.Now()
	sort.InsertionMergeSortFunc(A, func(a, b int) bool {
		return a < b
	})
	fmt.Println("elapsed for insertion merge sort:", time.Now().Sub(start))
}
