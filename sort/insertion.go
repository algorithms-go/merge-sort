package sort

// InsertionSortFunc uses the insertion sort algorithm and a sorting function
func InsertionSortFunc(A []int, fn sortFunc) []int {
	if fn == nil {
		fn = func(a, b int) bool {
			return a < b
		}
	}
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
