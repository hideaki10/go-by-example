package main

import "fmt"

func main() {

	nums1 := make([]int, 3)
	nums1[0] = 3
	nums1[1] = 1
	nums1[2] = 2

	nums2 := make([]int, 2)
	nums2[0] = 1
	nums2[1] = 1

	res := intersect(nums1, nums2)
	fmt.Println(res)
	//subnets(nums)//sdd(nums)//permute(nums)//	sortedArrayToBST(nums)numsss := 1generate(numsss)
}

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	hashMap := make(map[int]int)
	if len(nums1) > len(nums2) {
		result = compare(nums2, nums1, result, hashMap)
	} else {
		result = compare(nums1, nums2, result, hashMap)
	}
	return result
}
func compare(small []int, large []int, result []int, hashMap map[int]int) []int {
	for _, value := range large {
		if _, ok := hashMap[value]; ok {
			hashMap[value] += 1
		} else {
			hashMap[value] = 1
		}

	}
	for _, value := range small {
		if mapValue, ok := hashMap[value]; ok {
			if mapValue != 0 {
				hashMap[value] -= 1
				result = append(result, value)
			}
		}
	}
	return result
}
