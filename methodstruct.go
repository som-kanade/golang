package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func (p person) details() string {
	return string(p.name + "  " + p.city)
}

func main() {
	p1 := person{name: "andy", age: 28, city: "ac"}
	p2 := person{name: "john", age: 28, city: "de"}
	fmt.Println(p1)
	fmt.Println(p1.details())
	fmt.Println(p2.details())
}
