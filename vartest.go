package main

import "fmt"

var a int = 7

func test() {
	var a int = 5
	fmt.Println(a)
}

func main() {

	fmt.Println(a)
	test()

}
