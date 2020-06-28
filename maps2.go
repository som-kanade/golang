package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1)

	m2 := make(map[string]int)
	fmt.Println(m2)

	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m3)

	m4 := make(map[string]string)
	m4["john"] = "doe"
	m4["james"] = "danny"
	m4["jimmy"] = "alamia"

	fmt.Println(m4)
	fmt.Println(m4["james"])

	//_, ok := m4["john"]
	//fmt.Println(ok)

	delete(m4, "james")
	fmt.Println(m4)

	for i, v := range m4 {
		fmt.Println(i, v)
	}

}
