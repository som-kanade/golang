package main

import "fmt"

func main() {
	x := 20
	a := &x
	fmt.Println(x)
	fmt.Println(a)
	*a = 30
	fmt.Println(a)
}
