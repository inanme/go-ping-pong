package rndom

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func Test_split(t *testing.T) {
	split := func(output chan<- string, slice []string, done chan struct{}) {
		for _, e := range slice {
			output <- e
		}
		close(done)
	}

	done := make(chan struct{})
	output := make(chan string)

	go split(output, []string{"mert", "inan", "hello", "fd"}, done)

	for {
		select {
		case k := <-output:
			fmt.Println(k)
		case <-done:
			fmt.Println("here")
			close(output)
			return
		}
	}
}

func Test_stop(t *testing.T) {
	stop := func(stop chan struct{}, id int) {
		for {
			select {
			case <-stop:
				fmt.Printf("Stopping %d\n", id)
				return
			default:
				fmt.Printf("Waiting %d\n", id)
				time.Sleep(1 * time.Second)
			}
		}
	}
	dummyChannel := make(chan struct{})
	for i := 0; i < 10; i++ {
		go stop(dummyChannel, i)
	}

	time.Sleep(3 * time.Second)
	close(dummyChannel)
	time.Sleep(time.Second)
}

func Test_start(t *testing.T) {
	start := func(start chan struct{}, id int) {
		<-start
		fmt.Printf("mine is %d\n", id)
	}
	dummyChannel := make(chan struct{})
	for i := 0; i < 10; i++ {
		go start(dummyChannel, i)
	}

	time.Sleep(2 * time.Second)
	close(dummyChannel)
	time.Sleep(2 * time.Second)
}

func Test_Downloadurl(t *testing.T) {
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

type payload struct {
	result []string
}

func (p *payload) UnmarshalJSON(b []byte) error {
	var points []json.RawMessage
	if err := json.Unmarshal(b, &points); err != nil {
		return err
	}

	for _, point := range points {
		var local []json.RawMessage
		if err := json.Unmarshal(point, &local); err != nil {
			return err
		}
		var k int64
		if err := json.Unmarshal(local[0], &k); err != nil {
			return err
		}
		p.result = append(p.result, time.Unix(k, 0).Format(time.RFC3339))
	}
	return nil
}

func TestPayload(t *testing.T) {
	var s payload
	err := json.Unmarshal([]byte(`[[1648146930,"1"],[1648146945,"1"]]`), &s)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, []string{"2022-03-24T18:35:30Z", "2022-03-24T18:35:45Z"}, s.result)
}
