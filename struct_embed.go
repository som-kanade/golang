package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
	age       int
}

type secretagent struct {
	person
	lt bool
}

func main() {
	sa1 := secretagent{
		person: person{firstname: "john", lastname: "doe", age: 23},
		lt:     true,
	}

	fmt.Println(sa1.firstname)
}
