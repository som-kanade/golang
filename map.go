package main

import "fmt"

func main() {
	m := map[string]int{
		"john": 22,
		"doe":  24,
	}
	fmt.Println(m)

	m["henry"] = 24
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "doe")
	fmt.Println(m)

	for k1, v1 := range m {
		fmt.Println(k1, v1)
	}
}
