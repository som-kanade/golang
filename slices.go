package main

import "fmt"

func main() {
	x := make([]int, 5, 100)
	fmt.Println(x)
	x = append(x, 4, 5, 6)
	fmt.Println(x)
}
