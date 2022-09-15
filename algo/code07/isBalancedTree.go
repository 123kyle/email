package code07

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return process(root).IsB
}

type Info struct {
	IsB    bool
	Height int
}

func process(node *TreeNode) Info {
	if node == nil {
		return Info{true, 0}
	}
	leftInfo := process(node.Left)
	rightInfo := process(node.Right)
	isB := true
	if !leftInfo.IsB || !rightInfo.IsB {
		isB = false
	} else if abs(leftInfo.Height-rightInfo.Height) > 1 {
		isB = false
	}
	height := max(leftInfo.Height, rightInfo.Height) + 1
	return Info{isB, height}
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
