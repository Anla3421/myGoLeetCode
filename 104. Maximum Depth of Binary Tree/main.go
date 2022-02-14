package main

import (
	"fmt"
)

func main() {
	// leftNode := &TreeNode{}
	rightNode := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	input := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: rightNode,
	}
	result := maxDepth(input)
	fmt.Println("result", result)
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDepthSub(root, 0)
}

func maxDepthSub(root *TreeNode, i int) int {
	if root == nil {
		return i
	}
	left := maxDepthSub(root.Left, i+1)
	right := maxDepthSub(root.Right, i+1)
	return Max(left, right)
}

func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
