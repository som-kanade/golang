package main

import (
	"fmt"
)

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
	}
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("There")
	}

}

func main() {
	go foo()
	go bar()

}
