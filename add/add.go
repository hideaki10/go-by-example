package main

import "fmt"

//闭包
func adder() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}

}

func main() {
	a := adder()

	for i := 1; i < 10; i++ {
		fmt.Printf(" 0 + 1 + .... %d  = %d/n", i, a(i))
	}
}
