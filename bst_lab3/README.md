
# Binary Search Tree


#### Testing Insertion and Merge Sort

```sh
    $ go test ./bst_lab3/ -v
```

```
=== RUN   TestBSTAdd
--- PASS: TestBSTAdd (0.00s)
=== RUN   TestInOrder
--- PASS: TestInOrder (0.00s)
=== RUN   TestPostOrder
--- PASS: TestPostOrder (0.00s)
=== RUN   TestPreOrder
--- PASS: TestPreOrder (0.00s)
=== RUN   TestSearch
--- PASS: TestSearch (0.00s)
=== RUN   TestRemove
    bst_test.go:140: Expected '[10 5 1 8 52 30 45]' but got '[10 5 1 8 52 45 30]'
--- FAIL: TestRemove (0.00s)
=== RUN   TestSmallest
--- PASS: TestSmallest (0.00s)
=== RUN   TestLargest
--- PASS: TestLargest (0.00s)
FAIL
FAIL	github.com/orangeseeds/algoLabs/bst_lab3	0.009s
FAIL
```
