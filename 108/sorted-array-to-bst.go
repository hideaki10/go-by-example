package main

import "fmt"

func main() {

	nums := []int{-10, -3, 0, 5, 9}

	result := sortedArrayToBst(nums)

	fmt.Println(result)
}

func sortedArrayToBst(nums []int) *TreeNode {

	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	mid := (left + right) / 2

	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
