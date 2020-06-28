package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	p1 := person{
		firstname: "john",
		lastname:  "doe",
	}
	p2 := person{
		firstname: "shane",
		lastname:  "doe",
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstname)
	fmt.Println(p2.lastname)
}
