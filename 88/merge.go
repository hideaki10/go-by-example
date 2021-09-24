package main

import "fmt"

func main() {

	digits := make([]int, 6)

	digits[0] = 1
	digits[1] = 2
	digits[2] = 3
	digits[3] = 0
	digits[4] = 0
	digits[5] = 0
	digits2 := make([]int, 3)

	digits2[0] = 2
	digits2[1] = 5
	digits2[2] = 6

	m, n := 3, 3

	merge(digits, m, digits2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}

	}

	copy(nums1, sorted)
	fmt.Println(nums1)
}
