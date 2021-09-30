package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2, 3, 3, 4}
	majorityElement(nums)

}

func majorityElement(nums []int) int {
	tempMap := make(map[int]int)
	result := 0
	temp := 0

	for _, value := range nums {
		if _, ok := tempMap[value]; ok {
			tempMap[value] += 1
			continue
		}
		tempMap[value] = 1
	}

	for key, value := range tempMap {
		if value > temp {
			temp = value
			result = key
		}
	}
	fmt.Println(result)
	return result
}
