package main

import (
	"errors"
	"fmt"
)

type BST struct {
	root *Node
	size int
}

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Value  string
	Key    int
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) IsEmpty() bool {
	return bst.size == 0
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
		bst.root = newNode
		bst.size++
		return nil
	}

	parentNode := bst.root
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
			return fmt.Errorf("duplicate key %v not allowed.", key)
		}
	}
	return nil
}

func (bst *BST) Remove(key int) error {

	if bst.IsEmpty() {
		return errors.New("BST is empty")
	}

    bst.root = bst.removeNode(key, bst.root)
	return nil
}

func (bst *BST) Search(key int) (string, error) {
	if bst.IsEmpty() {
		return "", errors.New("BST is empty")
	}

	node := bst.root
	for node != nil {
		if key < node.Key {
			node = node.Left
		} else if key > node.Key {
			node = node.Right
		} else {
			return node.Value, nil
		}
	}
	return "", fmt.Errorf("key %v not found", key)
}

func (bst *BST) Smallest() (int, string, error) {
	if bst.IsEmpty() {
		return 0, "", fmt.Errorf("BST is empty")
	}

	node := bst.minNode(bst.root)
	return node.Key, node.Value, nil
}

func (bst *BST) Largest() (int, string, error) {
	if bst.IsEmpty() {
		return 0, "", fmt.Errorf("BST is empty")
	}

	node := bst.maxNode(bst.root)
	return node.Key, node.Value, nil
}

func (bst *BST) InOrderWalk() []int {
	walk := []int{}
	inOrder(bst.root, &walk)
	return walk
}

func (bst *BST) PreOrderWalk() []int {
	walk := []int{}
	preOrder(bst.root, &walk)
	return walk
}

func (bst *BST) PostOrderWalk() []int {
	walk := []int{}
	postOrder(bst.root, &walk)
	return walk
}

func (bst *BST) maxNode(root *Node) *Node {
	for root.Right != nil {
		root = root.Right
	}
	return root
}
func (bst *BST) minNode(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func (bst *BST) removeNode(key int, root *Node) *Node {
	if root == nil {
		return nil
	}

	if key < root.Key {
		root.Left = bst.removeNode(key, root.Left)
	} else if key > root.Key {
		root.Right = bst.removeNode(key, root.Right)
	} else {
		if root.Left != nil && root.Right != nil {
			min := bst.minNode(root.Right)
			root.Key = min.Key
			root.Right = bst.removeNode(min.Key, root.Right)
		} else {
			if root.Left == nil && root.Right == nil {
				root = nil
			} else if root.Left == nil {
				root = root.Right
			} else {
				root = root.Left
			}
			bst.size--
		}
	}
	return root
}

func inOrder(root *Node, walk *[]int) {
	if root != nil {
		inOrder(root.Left, walk)
		*walk = append(*walk, root.Key)
		inOrder(root.Right, walk)
	}
}

func preOrder(root *Node, walk *[]int) {
	if root != nil {
		*walk = append(*walk, root.Key)
		preOrder(root.Left, walk)
		preOrder(root.Right, walk)
	}
}

func postOrder(root *Node, walk *[]int) {
	if root != nil {
		postOrder(root.Left, walk)
		postOrder(root.Right, walk)
		*walk = append(*walk, root.Key)
	}
}
