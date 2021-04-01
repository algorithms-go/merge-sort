package sort

const insertionSortMaxThreshold = 20

type sortFunc func(a, b int) bool

// merge merges 2 slices into a single slice, also sorting the resulting slice
func merge(L, R []int, fn sortFunc) []int {
	A := make([]int, len(L)*2)
	i, j, k := 0, 0, 0
	for i < len(L) && j < len(R) {
		if fn(L[i], R[j]) {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
		k++
	}

	// check if any elements from the LEFT/RIGHT side were missed
	// in the comparison section above
	for i < len(L) {
		A[k] = L[i]
		i++
		k++
	}
	for j < len(R) {
		A[k] = R[j]
		j++
		k++
	}

	return A
}

// sort sorts the given slice by using the Merge Sort algorithm
// the algorithm uses the divide & conquer approach by dividing the problem into sub problems
// the algorithm works recursively, by calling itself
// HOW IT WORKS:
// 1. the slice is divided in 2 pieces (right, left)
// 2. the right and left sides are sorted first then they are merged
// 3. the sort function divides the incoming slice till it cannot be divided anymore (the slice has only 1 element)
// 4. after that it starts merging & sorting all the slices back
// 5. at the end we have 1 single slice sorted out
func sort(A []int, fn sortFunc) []int {
	if len(A) > 1 {
		m := len(A) / 2
		L := A[:m]
		R := A[m:]
		A = merge(sort(R, fn), sort(L, fn), fn)
	}
	return A
}

func hybridSort(A []int, fn sortFunc) []int {
	if len(A) > 1 {
		m := len(A) / 2
		if m <= insertionSortMaxThreshold {
			return insertionSortFunc(A, fn)
		}
		L := A[:m]
		R := A[m:]
		A = merge(hybridSort(R, fn), hybridSort(L, fn), fn)
	}
	return A
}

// MergeSort sorts a given slice by using the merge sorting algorithm
func MergeSort(A []int) []int {
	fn := func(a, b int) bool {
		return a < b
	}
	return sort(A, fn)
}

// MergeSort sorts a given slice by using the merge sorting algorithm and a sorting function
func MergeSortFunc(A []int, fn sortFunc) []int {
	return sort(A, fn)
}

// InsertionMergeSortFunc sorts a given slice by using a hybrid algorithm
// insertion & merge sorting algorithm and a sorting function
func InsertionMergeSortFunc(A []int, fn sortFunc) []int {
	return hybridSort(A, fn)
}
