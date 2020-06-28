package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {

	r := rect{1, 1}.area()
	fmt.Println(r)
}
