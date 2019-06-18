package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.amazon.co.uk",
	}
	var allDone sync.WaitGroup
	allDone.Add(len(urls))

	println(strings.Join(urls, "\n"))
	for _, url := range urls {
		go func(url string) {
			if response, err := http.Get(url); err == nil {
				if bytes, err := ioutil.ReadAll(response.Body); err == nil {
					fmt.Printf("the size of %s is %d\n", url, len(bytes))
				} else {
					fmt.Printf("Can not read %s\n", url)
				}
			} else {
				fmt.Printf("Can not get %s %s\n", url, err.Error())
			}
			allDone.Done()
		}(url)
	}
	allDone.Wait()
	fmt.Println("All done")
}
