package main

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// var (
// 	layer int
// 	left  int
// 	right int
// )

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return layer
// 	}
// 	if root.Left != nil {
// 		left = maxDepthSub(root.Left, left)
// 	}
// 	if root.Right != nil {
// 		right = maxDepthSub(root.Right, right)
// 	}
// 	fmt.Println("left", left)
// 	fmt.Println("right", right)
// 	max := Max(left, right)
// 	layer++
// 	return layer + max
// }

// func maxDepthSub(root *TreeNode, i int) int {
// 	if root == nil {
// 		return i
// 	}
// 	if root.Left != nil {
// 		left = maxDepthSub(root.Left, left)
// 	}
// 	if root.Right != nil {
// 		right = maxDepthSub(root.Right, right)
// 	}
// 	i++
// 	return i
// }

// func Max(a int, b int) int {
// 	if a >= b {
// 		return a
// 	}
// 	return b
// }
