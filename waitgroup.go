package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
	//wg.Done()
}

func main() {
	wg.Add(3)
	go say("hello")
	go say("world")
	go say("great!")
	wg.Wait()

}
