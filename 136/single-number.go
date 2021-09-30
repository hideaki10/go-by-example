package main

import "fmt"

func main() {
	nums := []int{1}

	singleNumbers(nums)
}

func singleNumbers(nums []int) int {

	result := 0
	tempMap := make(map[int]int)

	if len(nums) == 1 {
		result = nums[0]
		fmt.Println(result)
		return result
	}

	for _, value := range nums {
		if _, ok := tempMap[value]; ok {
			tempMap[value] += 1
			continue
		}
		tempMap[value] = 0
	}

	for key, value := range tempMap {
		if value == 0 {
			result = key
			break
		}
	}
	fmt.Println(result)
	return result
}
