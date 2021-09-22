package main

import "fmt"

func main() {

	nums := make([]int, 5)
	// nums[0] = -2
	// nums[1] = 1
	// nums[2] = -3
	// nums[3] = 4
	// nums[4] = -1
	// nums[5] = 2
	// nums[6] = 1
	// nums[7] = -7
	// nums[8] = 4

	nums[0] = 5
	nums[1] = 4
	nums[2] = -1
	nums[3] = 7
	nums[4] = 8

	// number := majorityElement(nums)
	// fmt.Println(number)

	maxSubArray(nums)

	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }

	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

}

func majorityElement(nums []int) int {
	var number, count int = 0, 0
	for _, num := range nums {
		if count == 0 {
			number = num
			fmt.Println("count == 0")
		}
		if number == num {
			fmt.Println("==num")
			count++
		} else {
			fmt.Println("!=num")
			count--
		}
	}
	return number
}

func maxSubArray(nums []int) int {

	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}

	}

	return maxSum

}

func twoSum(nums []int, targent int) []int {
	saveValueMap := make(map[int]int)
	result := []int{}

	for i, k := range nums {
		if value, exist := saveValueMap[targent-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}

		saveValueMap[k] = i
	}
	return result

}
