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
	var int100k []int
	for i := 0; i < n; i++ {
		int100k = append(int100k, rand.Intn(n))
	}
	start := time.Now()
	sort.InsertionSortFunc(int100k, func(a, b int) bool {
		return a < b
	})
	fmt.Println("elapsed", time.Now().Sub(start))

	start = time.Now()
	sort.MergeSortFunc(int100k, func(a, b int) bool {
		return a < b
	})
	fmt.Println("elapsed", time.Now().Sub(start))

	start = time.Now()
	sort.InsertionMergeSortFunc(int100k, func(a, b int) bool {
		return a < b
	})
	fmt.Println("elapsed", time.Now().Sub(start))
}
