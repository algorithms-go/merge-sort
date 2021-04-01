package sort

// insertionSortFunc uses the insertion sort algorithm and a sorting function
func insertionSortFunc(A []int, fn sortFunc) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i > -1 && fn(key, A[i]) {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}
	return A
}
