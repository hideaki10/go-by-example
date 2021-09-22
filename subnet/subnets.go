package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 5)

	nums[0] = -10
	nums[1] = -3
	nums[2] = 0
	nums[3] = 5
	nums[4] = 9
	//subnets(nums)
	//sdd(nums)
	//permute(nums)
	//	sortedArrayToBST(nums)
	numsss := 1
	generate(numsss)
}

func subnets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		fmt.Println("这是在第 ", i)
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		list = append(list, nums[i])

		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})

	return res
}

// 判断 有和没有 所以有2个递归
func sdd(nums []int) [][]int {
	res := [][]int{}

	var defs func(i int, list []int)
	defs = func(i int, list []int) {
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}

		// 这个递归是判断有的递归
		// list 储存 list + num【i】 比如 1 + 2  list【1，2】
		//
		list = append(list, nums[i])
		fmt.Println("我进第一个递归了", i)
		defs(i+1, list)

		// 这个递归是判断没有的递归
		// list 存储 没有的 就是夫亲的节点的值 比如 1 下面的节点 有 【1，2】 和 【1】 那这个递归就从右边的1开始
		//
		list = list[:len(list)-1]
		fmt.Println("我进第二个递归了", i)
		defs(i+1, list)
	}
	defs(0, []int{})

	return res
}

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var defs func(list []int)

	defs = func(list []int) {

		if len(list) == len(nums) {
			temp := make([]int, len(list))
			copy(temp, list)
			res = append(res, temp)
			return
		}

		for _, n := range nums {
			if visited[n] {
				continue
			}
			list = append(list, n)
			visited[n] = true
			defs(list)

			//往上退
			list = list[:len(list)-1] // [1,2,3] → [1,2]
			visited[n] = false        // 1
		}

	}

	defs([]int{})

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
