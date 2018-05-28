// https://www.youtube.com/watch?v=uBjoTxosSys
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^/(.+)\\.(.+)$")
	path := r.URL.Path
	matches := re.FindAllStringSubmatch(path, -1)
	if matches != nil {
		fmt.Fprintf(w, "hello there %s %s\n", matches[0][1], matches[0][2])
	} else {
		fmt.Fprint(w, "hello there")
	}
}
