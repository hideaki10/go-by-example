package main

import "fmt"

func main() {

	c := map[int]T{0: T{}}

	if project, ok := c[0]; ok {
		project.Number = 1
		c[0] = project
	}

	fmt.Println(c)

}

type T struct {
	Number int
	Text   string
}
