package main

import "fmt"

func add(num1 int, num2 int) (int, int) {
	var result1 = num1 + num2
	var result2 = num1 - num2
	return result1, result2
}

func main() {

	res1, res2 := add(5, 4)
	fmt.Println(res1, res2)

}
