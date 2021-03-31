package sort

func merge(L, R []int) []int {
	A := make([]int, len(L)*2)
	i, j, k := 0, 0, 0
	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
		k++
	}

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
func sort(A []int) []int {
	if len(A) > 1 {
		m := len(A) / 2
		L := A[:m]
		R := A[m:]
		A = merge(sort(R), sort(L))
	}
	return A
}

func MergeSort(A []int) []int {
	return sort(A)
}
