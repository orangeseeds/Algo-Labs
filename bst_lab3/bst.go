package main

import "errors"

type BST struct {
	Root *Node
	size int
}

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Value  string
	Key    int
}

func NewBST() BST {
	return BST{}
}

func (bst *BST) IsEmpty() bool {
	return bst.Size() == 0
}

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) Add(key int, value string) error {
	newNode := &Node{
		Key:   key,
		Value: value,
	}

	if bst.IsEmpty() {
		bst.Root = newNode
		bst.size++
		return nil
	}

	parentNode := bst.Root
	for {
		if newNode.Key < parentNode.Key {
			if parentNode.Left == nil {
				newNode.Parent = parentNode
				parentNode.Left = newNode
				bst.size++
				break
			}
			parentNode = parentNode.Left
		} else if newNode.Key > parentNode.Key {
			if parentNode.Right == nil {
				newNode.Parent = parentNode
				parentNode.Right = newNode
				bst.size++
				break
			}
			parentNode = parentNode.Right
		} else {
			return errors.New("Duplicate key not allowed.")
		}
	}

	return nil
}

func (bst *BST) Remove() {}

func (bst *BST) Search(key int) (string, error) {
	if bst.IsEmpty() {
		return "", errors.New("BST is empty")
	}

	node := bst.Root
	for node != nil {
		if key < node.Key {
			node = node.Left
		} else if key > node.Key {
			node = node.Right
		} else {
			return node.Value, nil
		}
	}
	return "", errors.New("Key not found")
}

func (bst *BST) Smallest() {}
func (bst *BST) Largest()  {}

func (bst *BST) InOrderWalk()   {}
func (bst *BST) PostOrderWalk() {}
