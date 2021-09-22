package main

import "fmt"

func main() {
	// var p People = LaoMiao{}
	// p.Eat("tao zi")
	Run(&LaoMiao{})
	Run(LaoSun{})
	Run(&GaoDan{})

	var people People
	people = LaoMiao{}
	val, ok := people.(LaoSun)
	if !ok {
		fmt.Println("dd")
	}

	fmt.Println(val)
}

type People interface {
	Eat(string) error
	Drink(string) error
}

type LaoMiao struct {
	Name string
	Age  int
}

func (l LaoMiao) Eat(thing string) error {
	fmt.Println("eat " + thing)
	return nil
}

func (l LaoMiao) Drink(thing string) error {
	fmt.Println("Drink " + thing)
	return nil
}

type LaoSun struct {
	Name string
	Age  int
}

func (l LaoSun) Eat(thing string) error {
	fmt.Println("is eating " + thing)
	return nil
}

func (l LaoSun) Drink(thing string) error {
	fmt.Println("is drinking " + thing)
	return nil
}

func Run(p People) {
	thing1, thing2 := "taozi", "taozi2"

	p.Eat(thing1)
	p.Drink(thing2)
}

type GaoDan struct {
	Name string
	Age  int
}

func (l *GaoDan) Eat(thing string) error {
	fmt.Println("你管我吃" + thing)
	return nil
}

func (l *GaoDan) Drink(thing string) error {
	fmt.Println("你管我喝" + thing)
	return nil
}

type Student interface {
	People
	Study()
}

type Empty interface{}

var str Empty = "zifuchuang"
var num Empty = 222
