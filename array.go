package main

import "fmt"

func main() {
	//a := [5]int{}
	//fmt.Println(a)
	//b := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(b)
	//c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	//fmt.Println(c)
	//fmt.Println(len(c))

	//d := [...]string{"avv", "bds", "cerr", "dth", "ehtggtg"}
	//for i := 0; i < len(d); i++ {
	//	fmt.Println(d[i])
	//}

	//e := [...]int{1, 2, 3, 4, 5, 6}
	//sum := 0
	//for i := 0; i < len(e); i++ {
	//	sum = sum + e[i]
	//}
	//fmt.Println(sum)

	//x := [...]int{1, 2, 3, 4, 5, 6}
	//for i, v := range x {
	//	fmt.Println("the index is %d  and thr value is %d ", i, v)
	//}

	y := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	sum := 0
	for _, v := range y {
		sum = sum + v
	}

	fmt.Println(sum)

}
