package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func setup() *BST {
	bst := NewBST()
	_ = bst.Add(10, "Value for 10")
	_ = bst.Add(52, "Value for 52")
	_ = bst.Add(5, "Value for 5")
	_ = bst.Add(8, "Value for 8")
	_ = bst.Add(1, "Value for 1")
	_ = bst.Add(40, "Value for 40")
	_ = bst.Add(30, "Value for 30")
	_ = bst.Add(45, "Value for 45")
	return bst
}

func TestBSTAdd(t *testing.T) {
	bsTree := NewBST()

	if bsTree.Size() != 0 {
		t.Errorf("BST size expected 0 but got %v", bsTree.Size())
	}

	_ = bsTree.Add(15, "Value is 15")
	if bsTree.Size() != 1 {
		t.Errorf("BST size expected 1 but got %v", bsTree.Size())
	}

	_ = bsTree.Add(10, "Value is 10")
	if bsTree.Size() != 2 {
		t.Errorf("BST size expected 2 but got %v", bsTree.Size())
	}

    result, _ := bsTree.Search(10)
	if result != "Value is 10" {
		t.Errorf("Expected 'Value is 10' but got %s", result)
	}

	result, _ = bsTree.Search(15)
	if result != "Value is 15" {
		t.Errorf("Expected 'Value is 15' but got %s", result)
	}
}

func TestInOrder(t *testing.T) {
	bst := setup()

	expected := []int{1, 5, 8, 10, 30, 40, 45, 52}
	got := bst.InOrderWalk()
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}

	expected = []int{1, 5, 8, 10, 25, 30, 40, 45, 52}
	_ = bst.Add(25, "Value for 25")
	got = bst.InOrderWalk()
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}

}

func TestPostOrder(t *testing.T) {
	bst := setup()

	expected := []int{1, 8, 5, 30, 45, 40, 52, 10}
	got := bst.PostOrderWalk()
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}

	expected = []int{1, 8, 5, 25, 30, 45, 40, 52, 10}
	_ = bst.Add(25, "Value for 25")
	got = bst.PostOrderWalk()
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestPreOrder(t *testing.T) {
	bst := setup()

	expected := []int{10, 5, 1, 8, 52, 40, 30, 45}
	got := bst.PreOrderWalk()
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}

	expected = []int{10, 5, 1, 8, 52, 40, 30, 25, 45}
	_ = bst.Add(25, "Value for 25")
	got = bst.PreOrderWalk()
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestSearch(t *testing.T) {
	bst := setup()

	expected := "Value for 40"
	got, _ := bst.Search(40)
	if expected != got {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}

	got, err := bst.Search(90)
	if err == nil {
		t.Errorf("expected 'key not found error' but got '%v'", got)
	}

	expected = "Value for 90"
	_ = bst.Add(90, "Value for 90")
	got, _ = bst.Search(90)
	if expected != got {
		t.Errorf("Expected '%s' but got '%s'", expected, got)
	}
}

func TestRemove(t *testing.T) {
    bst := setup()
    _ = bst.Remove(40)
    if bst.Size() != 7 {
        t.Errorf("expected 7 but got %v", bst.Size())
    }

    expected := []int{1, 5, 8, 10, 30, 45, 52}
    got := bst.InOrderWalk()
    if !slices.Equal(expected, got) {
		t.Errorf("Expected '%v' but got '%v'", expected, got)
    }

    expected = []int{10, 5, 1, 8, 52, 30, 45}
    got = bst.PreOrderWalk()
    if !slices.Equal(expected, got) {
		t.Errorf("Expected '%v' but got '%v'", expected, got)
    }

}

func TestSmallest(t *testing.T) {
	bst := setup()
	_ = bst.Add(6, "Value for 6")
	_ = bst.Add(4, "Value for 4")
	_ = bst.Add(0, "Value for 0")
	_ = bst.Add(32, "Value for 32")

	key, value, _ := bst.Smallest()
	if key != 0 || value != "Value for 0" {
		t.Errorf("expected 0, 'Value for 0' but got %v, '%v'", key, value)
	}
}

func TestLargest(t *testing.T) {
	bst := setup()
	_ = bst.Add(6, "Value for 6")
	_ = bst.Add(54, "Value for 54")
	_ = bst.Add(0, "Value for 0")
	_ = bst.Add(32, "Value for 32")

	key, value, _ := bst.Largest()
	if key != 54 || value != "Value for 54" {
		t.Errorf("expected 54, 'Value for 54' but got %v, '%v'", key, value)
	}
}
