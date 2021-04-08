# Merge Sort Algorithm

<img alt="thumbnail image" src="https://github.com/algorithms-go/merge-sort/blob/master/thumbnail.jpg?raw=true" width="800"/>

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

<img alt="from js to go image" src="https://github.com/algorithms-go/merge-sort/blob/master/gifs/merge-sort-cards-intro.gif?raw=true" width="800"/> 

### Overview

<img alt="algorithm overview image" src="https://github.com/algorithms-go/merge-sort/blob/master/gifs/algorithm-overview.gif?raw=true" width="800"/>

### Divide & Conquer principle

<img alt="divide and conquer image" src="https://github.com/algorithms-go/merge-sort/blob/master/gifs/divide-and-conquer.gif?raw=true" width="800"/>

### The Algorithm

### `merge` function

<img alt="merge function image" src="https://github.com/algorithms-go/merge-sort/blob/master/gifs/merge-function.gif?raw=true" width="800"/>

### `merge` function pseudocode

```
function merge(L, R) {
  let A[0..L.length+R.length]
  i, j, k = 0, 0, 0
  // compare LEFT & RIGHT side elements before merging
  while i < L.length and j < R.length {
    if L[i] < R[j] {
      A[k] = L[i]
      i++
    } else {
      A[k] = R[j]
      j++
    }
  }
  
  // check if any elements from the left/right side
  // were missed in the comparison section above
  while i < L.length {
    A[k] = L[i]
    i++
    k++
  }
  while j < R.length {
    A[k] = R[j]
    j++
    k++
  }
  
  return A
}
```

As you can see from the pseudocode above we’re mainly doing 2 things:

1. **Compare** the elements and decide which ones goes first in the resulting **A** array/slice.

```
...
while i < L.length and j < R.length {
  if L[i] < R[j] {
    A[k] = L[i]
    i++
  } else {
    A[k] = R[j]
    j++
  }
}
...
```

Also, notice the **sorting condition** above is hardcoded.
Instead, we could also receive a **callback function** that accepts `a` and `b`
and returns `bool`, giving the control of **sorting** to the caller.

Something like the following:

```type sortFunc func(a, b int) bool```

2. At the end we check for any left elements and we add them at the end of
the resulting array/slice A

```
...
while i < L.length {
  A[k] = L[i]
  i++
  k++
}
while j < R.length {
  A[k] = R[j]
  j++
  k++
}
return A
...
```

For the Go implementation check out the [merge function](sort/merge.go#L11)


### `sort` function

<img alt="sort function image" src="https://github.com/algorithms-go/merge-sort/blob/master/gifs/sort-function.gif?raw=true" width="800"/>

### `sort` function pseudocode

```
function sort(A) {
  if A.length > 1 {
    n1 = A.length / 2
    n2 = A.length - n1
    let L[0..n1] and R[0..n2]
    
    // split A in half => L and R arrays/slices
    for i=0 to n1 {
      L[i] = A[i]
    }
    for j=0 to n2 {
      R[j] = A[n1+j]
    }
    // for languages that use clever constructs
    // m = A.length / 2
    // L = A[:m]
    // R = A[m:]
    // recursively sort & merge L & R arrays/slices
    // A starts growing when recursion reaches the very bottom
    A = merge(sort(L), sort(R))
  }
  
  // return when array/slice consists of 1 element
  return A
}
```

The sort function primarily consists of 3 steps:
1. **Split** the incoming array/slice in **2 sides**, on every recursive call
& create *LEFT* & *RIGHT* arrays/slices.

```
...
n1 = A.length / 2
n2 = A.length - n1
let L[0..n1] and R[0..n2]
    
// split A in half => L and R arrays/slices
for i=0 to n1 {
  L[i] = A[i]
}
for j=0 to n2 {
  R[j] = A[n1+j]
}
...
```

2. **Return** array/slice when it *cannot be divided anymore*, or it has **only 1 element**.

```
if A.length > 1 {
  ...
}
// Here we have an array/slice of 1 element only
return A
```

3. **Call sort and merge** for **LEFT** & **RIGHT** sides **recursively**, till exit condition
is met (*array/slice has 1 element only or both sides were merged*).

```
...
// recursively sort & merge L & R arrays/slices
// A starts growing when recursion reaches the very bottom
A = merge(sort(L), sort(R))
...
```

For the Go implementation check out the [sort function](sort/merge.go#L51)
and [insertionMergeSort function](sort/merge.go#L51)

### Facts

As I say in everyone of my tutorials, there is really ***no perfect tool/solution*** or algorithm that does it all.
You really have to use the right tool for the right problem, or I should say use the right
Sorting Algorithm for the right situation. That’s why before we wrap this up, let’s explore
couple of ***facts*** about the ***Merge Sort Algorithm***, so that you know when is best to use it and when not:

- Merge Sort is ***Efficient*** & ***Very Fast*** for ***big data*** sets
- Merge Sort Time Complexity is: `O(n*log(n))`
- Merge Sort Space Complexity is: `O(n)`

Here are couple of facts about the Insertion Sort Algorithm:

- Insertion Sort is ***Efficient*** for ***small data*** sets
- Insertion Sort ***Best*** Time Complexity is: `O(n)`
- Insertion Sort ***Worst/Average*** Time Complexity is: `O(n*n)`
- Insertion Sort Space Complexity is: `O(1)`

Knowing the above facts, helps ***identify & diagnose*** the sorting problem and
give us an idea of how we shall proceed with sorting and what kind of sorting algorithms to use.

### How to Test

```bash
# cd into the package
cd sort

# run benchmark tests for BenchmarkInsertionMergeSortFunc
# feel free to change the above benchmark name with your own
go test -v -bench 'BenchmarkInsertionMergeSortFunc'
```

### Benchmarks

The current benchmarks ran on a relatively small
data set of **100 000** randomly generated integers.
In real life, data sets can be *bigger*, and the *distribution*
of the data can be **worse** in terms of how sorted it is.

Thus, as you can see the ***Merge Sort Algorithm***, as well as the
hybrid algorithm using the ***Insertion Sort Algorithm***,
and the merge algorithm perform way better, and they are
pretty ***stable algorithms***, which makes the result ***very predictable***.

The results may vary, due to the randomly generated numbers,
as well as depending on the specific environment and machine
they are run on.

These are the results I got on my personal Macbook Pro 2019 machine,
yours might be different:

Bottom line, the indicating factor that one algorithm performs
the best is how many time it can be run in a time window
& how many ns/op it takes.

Pay attention at the **1st number** which is how many times the function was executed
in a small time window (*bigger is better*), also at the **2nd number** which
shows how many nanoseconds/operation it took (*smaller is better*).

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

- [Merge Sort Algorithm - Medium Article](https://steevehook.medium.com/merge-sorting-algorithm-in-go-f7d5b47e015b)
- [Merge Sort - Geeks for Geeks](https://www.geeksforgeeks.org/merge-sort/)
- [Merge Sort - Programiz](https://www.programiz.com/dsa/merge-sort)
- [Big O Cheatsheet - Insertion vs Merge vs Quick sort](https://www.bigocheatsheet.com/)
- [Merge Sort - Wiki](https://en.wikipedia.org/wiki/Merge_sort)
- [Insertion Merge Sort - Ford–Johnson algorithm - Wiki](https://en.wikipedia.org/wiki/Merge-insertion_sort)
- [Divide & Conquer Principle - Wiki](https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm)
