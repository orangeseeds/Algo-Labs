package main

import (
	"testing"
)

func assertEqual[K comparable](t *testing.T, result K, expected K) {
	if result != expected {
		t.Errorf("expected %v but got %v", expected, result)
	}}

var bst BST

func init() {
	bst = NewBST()
	bst.Add(10, "value for 10")
	bst.Add(52, "value for 52")
	bst.Add(5, "value for 5")
	bst.Add(8, "value for 8")
	bst.Add(1, "value for 1")
	bst.Add(40, "value for 40")
	bst.Add(30, "value for 30")
	bst.Add(45, "Value for 45")
}

func TestBSTAdd(t *testing.T) {
	bsTree := NewBST()

	// if bsTree.Size() != 0 {
	// 	t.Errorf("BST size expected 0 but got %v", bsTree.Size())
	// }
    assertEqual(t, 0, bsTree.Size())

	_ = bsTree.Add(15, "Value is 15")
	// if bsTree.Size() != 1 {
	// 	t.Errorf("BST size expected 1 but got %v", bsTree.Size())
	// }

	_ = bsTree.Add(10, "Value is 10")
	if bsTree.Size() != 2 {
		t.Errorf("BST size expected 2 but got %v", bsTree.Size())
	}

	res, _ := bsTree.Search(15)
	if res != "Value is 15" {
		t.Errorf("Expected 'Value is 15' but got %s", res)
	}

	res, _ = bsTree.Search(10)
	if res != "Value is 10" {
		t.Errorf("Expected 'Value is 10' but got %s", res)
	}
}

func TestInOrder(t *testing.T) {}

func TestPostOrder(t *testing.T) {}

func TestPreOrder(t *testing.T) {}

func TestSearch(t *testing.T) {
	expected := "value for 40"
	got, err := bst.Search(40)
	if err != nil {
		t.Error(err)
	}
	if expected != got {
		t.Errorf("Expected %s but got %s", expected, got)
	}

	expected = "Value for 90"
	_ = bst.Add(90, "Value for 90")
	got, err = bst.Search(90)
	if err != nil {
		t.Error(err)
	}

	if expected != got {
		t.Errorf("Expected %s but got %s", expected, got)
	}
}

func TestRemove(t *testing.T) {
}

func TestSmallest(t *testing.T) {}

func TestLargest(t *testing.T) {}
