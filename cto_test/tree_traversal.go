package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	output := make([]int, 0)

	output = append(output, root.Val)
	output = append(output, left...)
	output = append(output, right...)
	return output
}
