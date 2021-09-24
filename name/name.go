package main

import "fmt"

func main() {

	result := Sum(1)
	fmt.Println(result)

	s := new([]int)

	fmt.Println(s)

	r := []int{}

	fmt.Println(r)

	m := make([]int, 10)

	fmt.Println(m)

}

func Sum(num int) (newNum int) {

	newNum += 1
	fmt.Println(" 関数内", num)
	return
}
