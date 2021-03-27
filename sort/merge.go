package sort

import (
	"math"
)

// merge merges 2 sorted slices & sorts them into 1 final slice
// Note: we work with []int, because numbers are the easiest to reason about
func merge(A []int, p, q, r int) []int {
	n1, n2 := q-p+1, r-q
	L, R := make([]int, n1+1), make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = A[q+j+1]
	}

	i, j := 0, 0
	L[n1], R[n2] = math.MaxInt64, math.MaxInt64
	for k := p; k < r+1; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
	return A
}

// MergeSort sorts the given slice by using the Merge Sort algorithm
func MergeSort(A []int, p, q, r int) []int {
	return merge(A, p, q, r)
}
