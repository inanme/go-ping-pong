package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type getPageSizeResponse struct {
	id    int
	url   string
	size  int
	error string
}

func slave(id int, input chan string, output chan getPageSizeResponse) {
	url := <-input
	resp, err := http.Get(url)
	if err != nil {
		output <- getPageSizeResponse{id, url, 0, err.Error()}
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		output <- getPageSizeResponse{id, url, 0, err.Error()}
		return
	}
	output <- getPageSizeResponse{id, url, len(body), ""}
}

func main() {
	urls := [...]string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.bbc.co.uk"}
	var semaphore sync.WaitGroup
	semaphore.Add(len(urls))
	input := make(chan string)
	output := make(chan getPageSizeResponse)
	defer close(input)
	defer close(output)
	for i := 0; i < 10; i++ {
		go slave(i, input, output)
	}
	for _, e := range urls {
		input <- e
	}
	for i := 0; i < len(urls); i++ {
		resp := <-output
		fmt.Println(resp)
	}

}
