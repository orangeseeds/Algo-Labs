
# Insertion Sort and Merge Sort

[link to the graph](https://docs.google.com/spreadsheets/d/1DARufftgtmJ-bZV139z3gfqCX-kCkiBaFpmXTdDEq3w/edit#gid=1449020709)

**Merge Sort**  
  - Worst Case: O(n * log n)

**Insertion Sort**
  - Worst Case: O(n^2)

#### Testing Insertion and Merge Sort

```sh
    $ go test ./sort_lab2/ -bench MergeSortWorst
```

```
=== RUN   TestInsertionSort
--- PASS: TestInsertionSort (0.00s)
=== RUN   TestMergeSort
--- PASS: TestMergeSort (0.00s)
PASS
ok  	github.com/orangeseeds/algoLabs/sort_lab2	0.014s
  
```

#### Benchmarking Insertion Sort 

##### Worst Case
```sh
    $ go test ./sort_lab2/ -bench InsertionSortWorst
```

```
goos: linux
goarch: amd64
pkg: github.com/orangeseeds/algoLabs/sort_lab2
cpu: Intel(R) Pentium(R) CPU  N3530  @ 2.16GHz
BenchmarkInsertionSortWorst/array_size_1-4         	771637957	         1.616 ns/op
BenchmarkInsertionSortWorst/array_size_11-4        	 6400743	       159.6 ns/op
BenchmarkInsertionSortWorst/array_size_21-4        	 2574297	       480.2 ns/op
BenchmarkInsertionSortWorst/array_size_31-4        	 1000000	      1015 ns/op
BenchmarkInsertionSortWorst/array_size_41-4        	  717453	      1650 ns/op
BenchmarkInsertionSortWorst/array_size_51-4        	  509460	      2346 ns/op
BenchmarkInsertionSortWorst/array_size_61-4        	  365815	      3682 ns/op
BenchmarkInsertionSortWorst/array_size_71-4        	  271370	      4477 ns/op
BenchmarkInsertionSortWorst/array_size_81-4        	  209758	      5801 ns/op
BenchmarkInsertionSortWorst/array_size_91-4        	  171169	      7238 ns/op
```

#### Benchmarking Merge Sort

##### Worst Case

```sh
    $ go test ./sort_lab2/ -bench MergeSortWorst
```

```
goos: linux
goarch: amd64
pkg: github.com/orangeseeds/algoLabs/sort_lab2
cpu: Intel(R) Pentium(R) CPU  N3530  @ 2.16GHz
BenchmarkMergeSortWorst/array_size_1-4         	206545437	         5.937 ns/op
BenchmarkMergeSortWorst/array_size_11-4        	  695575	      3217 ns/op
BenchmarkMergeSortWorst/array_size_21-4        	  275106	      6395 ns/op
BenchmarkMergeSortWorst/array_size_31-4        	  123033	     10815 ns/op
BenchmarkMergeSortWorst/array_size_41-4        	  158704	     15483 ns/op
BenchmarkMergeSortWorst/array_size_51-4        	   57633	     20020 ns/op
BenchmarkMergeSortWorst/array_size_61-4        	   92475	     22698 ns/op
BenchmarkMergeSortWorst/array_size_71-4        	   37402	     29173 ns/op
BenchmarkMergeSortWorst/array_size_81-4        	   35802	     28266 ns/op
BenchmarkMergeSortWorst/array_size_91-4        	   49194	     39080 ns/op
```
