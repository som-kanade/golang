package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func sendRequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d %s \n", res.StatusCode, url)

}

func main() {
	if (len(os.Args)) < 2 {
		fmt.Println("please pass arguments")
	}
	for _, url := range os.Args[1:] {
		go sendRequest("https://" + url)

	}
}
