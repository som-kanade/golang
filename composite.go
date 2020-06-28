package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 5, 5}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	for j := 0; j < 4; j++ {
		fmt.Println(j, x[j])
	}

}
