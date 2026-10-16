package main

import "fmt"

type TreeNode[T any] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

type BinaryTree[T any] struct {
	root *TreeNode[T]
}

// InsertHelper builds the subtree interactively and returns its root (nil if the user says "n").
func InsertHelper[T any](tree *TreeNode[T]) *TreeNode[T] {
	var exists string
	fmt.Print("Does node exist? (y/n): ")
	fmt.Scan(&exists)

	if exists != "y" {
		return nil
	}

	var data T
	fmt.Print("Enter data: ")
	fmt.Scan(&data)

	if tree == nil {
		tree = &TreeNode[T]{}
	}
	tree.data = data

	fmt.Printf("-- left of %v --\n", data)
	tree.left = InsertHelper(tree.left)

	fmt.Printf("-- right of %v --\n", data)
	tree.right = InsertHelper(tree.right)

	return tree
}

func (t *BinaryTree[T]) Insert() {
	t.root = InsertHelper(t.root) // the assignment now sticks, because t is a pointer to the wrapper
}

func PrintTreeHelper[T any](tree *TreeNode[T]) {
	if tree == nil {
		return
	}

	fmt.Print(tree.data, ": ")
	if tree.left != nil {
		fmt.Print("L -> ", tree.left.data, "  ")
	}
	if tree.right != nil {
		fmt.Print("R -> ", tree.right.data)
	}
	fmt.Println()

	PrintTreeHelper(tree.left)
	PrintTreeHelper(tree.right)
}

func (t *BinaryTree[T]) PrintTree() {
	if t.root == nil {
		fmt.Println("(empty tree)")
		return
	}
	PrintTreeHelper(t.root)
}

func main() {
	fmt.Println("***** Binary Tree Implementation on Go *****")
	tree := &BinaryTree[int]{}

	tree.Insert()
	tree.PrintTree()
}
