package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testCase struct {
	name string
	in   []int
	out  []int
	fn   sortFunc
}

type mergeSortSuite struct {
	suite.Suite
}

func (s mergeSortSuite) Test_MergeSort() {
	cases := []testCase{
		{
			name: "ascending sorting func",
			in:   []int{8, 7, 6, 5, 4, 3, 2, 1},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			fn:   nil,
		},
		{
			name: "sorted array",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			fn:   nil,
		},
		{
			name: "reverse sorted array",
			in:   []int{8, 7, 6, 5, 4, 3, 2, 1},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			fn:   nil,
		},
		{
			name: "sorted halves",
			in:   []int{2, 4, 5, 7, 1, 2, 3, 6},
			out:  []int{1, 2, 2, 3, 4, 5, 6, 7},
			fn:   nil,
		},
	}

	for _, c := range cases {
		s.Run(c.name, func() {
			out := MergeSort(c.in)
			s.Equal(c.out, out)
		})
	}
}

func (s mergeSortSuite) Test_MergeSortFunc() {
	cases := []testCase{
		{
			name: "ascending sorting func",
			in:   []int{8, 7, 6, 5, 4, 3, 2, 1},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			fn: func(a, b int) bool {
				return a < b
			},
		},
		{
			name: "nil sorting func",
			in:   []int{8, 7, 6, 5, 4, 3, 2, 1},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			fn:   nil,
		},
		{
			name: "descending sorting func",
			in:   []int{6, 7, 4, 1, 2, 3, 8, 5},
			out:  []int{8, 7, 6, 5, 4, 3, 2, 1},
			fn: func(a, b int) bool {
				return a > b
			},
		},
		{
			name: "sorted array",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			out:  []int{8, 7, 6, 5, 4, 3, 2, 1},
			fn: func(a, b int) bool {
				return a > b
			},
		},
		{
			name: "sorted halves",
			in:   []int{2, 4, 5, 7, 1, 2, 3, 6},
			out:  []int{1, 2, 2, 3, 4, 5, 6, 7},
			fn:   nil,
		},
	}

	for _, c := range cases {
		s.Run(c.name, func() {
			out := MergeSortFunc(c.in, c.fn)
			s.Equal(c.out, out)
		})
	}
}

func (s mergeSortSuite) Test_InsertionMergeSortFunc() {
	cases := []testCase{
		{
			name: "use insertion sort",
			in:   []int{8, 7, 6, 5, 4, 3, 2, 1},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			fn: func(a, b int) bool {
				return a < b
			},
		},
		{
			name: "use insertion & merge sort",
			in:   []int{21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			out:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			fn: func(a, b int) bool {
				return a < b
			},
		},
	}

	for _, c := range cases {
		s.Run(c.name, func() {
			out := InsertionMergeSortFunc(c.in, c.fn)
			s.Equal(c.out, out)
		})
	}
}

func TestMergeSort(t *testing.T) {
	suite.Run(t, new(mergeSortSuite))
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(generateIntegers(10_000))
	}
}

func BenchmarkMergeSortFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortFunc(generateIntegers(10_000), func(a, b int) bool {
			return a < b
		})
	}
}

func BenchmarkInsertionMergeSortFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionMergeSortFunc(generateIntegers(10_000), func(a, b int) bool {
			return a < b
		})
	}
}

// DISCLAIMER: This takes a long time. Have patience before you run it
func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSortFunc(generateIntegers(10_000), func(a, b int) bool {
			return a < b
		})
	}
}

func generateIntegers(n int) []int {
	var A []int
	for i := 0; i < n; i++ {
		A = append(A, rand.Intn(n))
	}
	return A
}
