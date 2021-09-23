package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go fun1()
	go fun2()
	fmt.Println("zhu zhu zhu zhu zhu")
	wg.Wait()
	fmt.Println("zuse jiese wan¥bi")
}

func fun1() {
	for i := 1; i < 10; i++ {
		fmt.Println("IIIIIII", i)
	}
	wg.Done()
}

func fun2() {
	defer wg.Done()
	for j := 1; j < 10; j++ {
		fmt.Println("JJJJJJJJ", j)
	}

}
