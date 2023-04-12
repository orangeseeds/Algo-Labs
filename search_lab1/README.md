# Binary and Linear Search

[link to the graph](https://docs.google.com/spreadsheets/d/1DARufftgtmJ-bZV139z3gfqCX-kCkiBaFpmXTdDEq3w/edit?usp=sharing)

**Binary Search**  
  - Best Case: O(1)
  - Worst Case: O(log n)

**Linear Search**
  - Best Case: O(1)
  - Worst Case: O(n)

#### Testing Binary and Linear Search

```sh
  $ go test ./search_lab1/ -v
```

```
=== RUN   TestLinearSearch
--- PASS: TestLinearSearch (0.00s)
=== RUN   TestBinarySearch
--- PASS: TestBinarySearch (0.00s)
PASS
ok  	github.com/orangeseeds/algoLabs/search_lab1	0.005s
  
```

#### Benchmarking Binary Search

##### Worst Case
```sh
  $ go test ./search_lab1/ -bench BinarySearchWorst
```

```
goos: linux
goarch: amd64
pkg: github.com/orangeseeds/algoLabs/search_lab1
cpu: Intel(R) Pentium(R) CPU  N3530  @ 2.16GHz
BenchmarkBinarySearchWorst/array_size_1000-4         	22423279	        50.38 ns/op
BenchmarkBinarySearchWorst/array_size_1100-4         	19891094	        58.04 ns/op
BenchmarkBinarySearchWorst/array_size_1200-4         	21798912	        54.04 ns/op
BenchmarkBinarySearchWorst/array_size_1300-4         	19808906	        54.02 ns/op
BenchmarkBinarySearchWorst/array_size_1400-4         	22229769	        57.71 ns/op
BenchmarkBinarySearchWorst/array_size_1500-4         	20745133	        54.40 ns/op
BenchmarkBinarySearchWorst/array_size_1600-4         	22001120	        54.13 ns/op
BenchmarkBinarySearchWorst/array_size_1700-4         	21992220	        56.05 ns/op
BenchmarkBinarySearchWorst/array_size_1800-4         	22190576	        54.41 ns/op
BenchmarkBinarySearchWorst/array_size_1900-4         	22284235	        56.07 ns/op
```

##### Best Case
```sh
  $ go test ./search_lab1/ -bench BinarySearchBest
```

```
goos: linux
goarch: amd64
pkg: github.com/orangeseeds/algoLabs/search_lab1
cpu: Intel(R) Pentium(R) CPU  N3530  @ 2.16GHz
BenchmarkBinarySearchBest/array_size_1000-4         	28535079	        45.14 ns/op
BenchmarkBinarySearchBest/array_size_1100-4         	23433609	        45.03 ns/op
BenchmarkBinarySearchBest/array_size_1200-4         	26774292	        45.37 ns/op
BenchmarkBinarySearchBest/array_size_1300-4         	26802680	        45.07 ns/op
BenchmarkBinarySearchBest/array_size_1400-4         	22290555	        47.14 ns/op
BenchmarkBinarySearchBest/array_size_1500-4         	24417746	        49.30 ns/op
BenchmarkBinarySearchBest/array_size_1600-4         	26240017	        48.38 ns/op
BenchmarkBinarySearchBest/array_size_1700-4         	25577136	        45.49 ns/op
BenchmarkBinarySearchBest/array_size_1800-4         	26805237	        48.63 ns/op
BenchmarkBinarySearchBest/array_size_1900-4         	25069814	        47.68 ns/op
```

#### Benchmarking Linear Search

##### Worst Case

```sh
  $ go test ./search_lab1/ -bench LinearSearchWorst
```

```
goos: linux
goarch: amd64
pkg: github.com/orangeseeds/algoLabs/search_lab1
cpu: Intel(R) Pentium(R) CPU  N3530  @ 2.16GHz
BenchmarkLinearSearchWorst/array_size_1000-4         	  806394	      1572 ns/op
BenchmarkLinearSearchWorst/array_size_1100-4         	  737738	      1862 ns/op
BenchmarkLinearSearchWorst/array_size_1200-4         	  679738	      1885 ns/op
BenchmarkLinearSearchWorst/array_size_1300-4         	  551808	      2027 ns/op
BenchmarkLinearSearchWorst/array_size_1400-4         	  548850	      2210 ns/op
BenchmarkLinearSearchWorst/array_size_1500-4         	  410214	      2525 ns/op
BenchmarkLinearSearchWorst/array_size_1600-4         	  482852	      2792 ns/op
BenchmarkLinearSearchWorst/array_size_1700-4         	  440488	      2670 ns/op
BenchmarkLinearSearchWorst/array_size_1800-4         	  467929	      2997 ns/op
BenchmarkLinearSearchWorst/array_size_1900-4         	  447241	      3245 ns/op
```
##### Best Case

```sh
  $ go test ./search_lab1/ -bench LinearSearchBest
```

```
goos: linux
goarch: amd64
pkg: github.com/orangeseeds/algoLabs/search_lab1
cpu: Intel(R) Pentium(R) CPU  N3530  @ 2.16GHz
BenchmarkLinearSearchBest/array_size_1000-4         	383273588	         3.260 ns/op
BenchmarkLinearSearchBest/array_size_1100-4         	371141149	         3.373 ns/op
BenchmarkLinearSearchBest/array_size_1200-4         	354448944	         3.281 ns/op
BenchmarkLinearSearchBest/array_size_1300-4         	385969618	         3.363 ns/op
BenchmarkLinearSearchBest/array_size_1400-4         	320483283	         3.159 ns/op
BenchmarkLinearSearchBest/array_size_1500-4         	366456932	         3.442 ns/op
BenchmarkLinearSearchBest/array_size_1600-4         	366279939	         3.129 ns/op
BenchmarkLinearSearchBest/array_size_1700-4         	380406732	         3.272 ns/op
BenchmarkLinearSearchBest/array_size_1800-4         	329397606	         3.204 ns/op
BenchmarkLinearSearchBest/array_size_1900-4         	319056184	         3.298 ns/op
```
