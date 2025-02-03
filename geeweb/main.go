package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for l, i := range req.Header {
			fmt.Fprintf(w, "header[%q] = %q\n", l, i)
		}
	})
	r.Run(":9999")
}
