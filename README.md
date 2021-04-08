# Merge Sort Algorithm

<img alt="from js to go image" src="https://github.com/algorithms-go/merge-sort/blob/master/thumbnail.jpg?raw=true" width="700"/>

The merge sort algorithm may seem complicated to implement because it uses a recursive approach,
which bugs a lot of people. The easiest way to reason about the merge sort algorithm is:
Imagine you have a big pile of cards that you want to sort out.

The principle is simple, divide the pile of cards in 2 pieces (left & right) till
new smaller piles consist of one card only. Then start merging the left & right piles
till you reach the top of the pile.

When merging 2 piles of cards compare each card with the other card in the pile and
merge them in the correct (sorted) order.

At the end of this recursive process we shall have the huge pile of cards sorted out.

To make what was described even clearer check out the following example with cards:

<img alt="from js to go image" src="https://github.com/algorithms-go/merge-sort/blob/master/gifs/merge-sort-cards-intro.gif?raw=true" width="700"/> 

### Overview

### The Algorithm

### Pseudocode

### Facts

### How to Test

```bash
# run benchmark tests for BenchmarkInsertionMergeSortFunc
go test -v -bench 'BenchmarkInsertionMergeSortFunc'
```

### Benchmarks

The current benchmarks ran on a relatively small
data set of 100 000 randomly generated integers.
In real life, data sets can be bigger, and the distribution
of the data can be worse in terms of how sorted it is.

Thus, as you can see the merge sorting algorithm, as well as the
hybrid algorithm using the insertion sorting algorithm,
and the merge algorithm perform way better, and they are
pretty stable algorithms, which makes the result very predictable.

The results may vary, due to the randomly generated numbers,
as well as depending on the specific environment and machine
they are run on.

Bottom line, the indicating factor that one algorithm performs
the best is how many time it can be run in a time window
& how many ns/op it takes.

```text
goos: darwin
goarch: amd64
pkg: github.com/algorithms-go/sorting/sort
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkInsertionSort
BenchmarkInsertionSort-16    	             21	           52535921 ns/op
PASS

goos: darwin
goarch: amd64
pkg: github.com/algorithms-go/sorting/sort
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMergeSort
BenchmarkMergeSort-16    	             747	   1608912 ns/op
PASS

goos: darwin
goarch: amd64
pkg: github.com/algorithms-go/sorting/sort
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMergeSortFunc
BenchmarkMergeSortFunc-16    	             705	   1614946 ns/op
PASS

goos: darwin
goarch: amd64
pkg: github.com/algorithms-go/sorting/sort
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkInsertionMergeSortFunc
BenchmarkInsertionMergeSortFunc-16    	     915	   1293846 ns/op
```

### Resources

- [Merge Sort - Geeks for Geeks](https://www.geeksforgeeks.org/merge-sort/)
- [Merge Sort - Programiz](https://www.programiz.com/dsa/merge-sort)
- [Merge Sort - Wiki](https://en.wikipedia.org/wiki/Merge_sort)
- [Insertion Merge Sort - Ford–Johnson algorithm - Wiki](https://en.wikipedia.org/wiki/Merge-insertion_sort)
- [Big O Cheatsheet - Insertion vs Merge vs Quick sort](https://www.bigocheatsheet.com/)
- [Divide & Conquer Principle](https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm)
