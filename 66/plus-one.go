package main

import "fmt"

func main() {

	digits := make([]int, 3)

	digits[0] = 1
	digits[1] = 2
	digits[2] = 3
	result := plusOne(digits)
	fmt.Println(result)
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 10 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}

	if digits[0] == 0 {

		newResult := make([]int, len(digits)+1)
		newResult[0] = 1

		for i := 1; i < len(digits); i++ {
			newResult[i] = digits[i]
		}

		return newResult

	} else {
		return digits
	}
}
