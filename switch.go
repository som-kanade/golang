package main

import "fmt"

func main() {
	switch {
	case (1 == 2):
		fmt.Println("not true")
	case (1 == 3):
		{
			fmt.Println("True")
		}
	default:
		fmt.Println("Cool case ")
	}
}
