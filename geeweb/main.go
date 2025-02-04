package main

import (
	"fmt"
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "hello world")
	})
	r.GET("/hello", func(c *gee.Context) {
		for l, i := range c.Req.Header {
			fmt.Fprintf(c.Writer, "header[%q] = %q\n", l, i)
		}
	})
	r.Run(":9999")
}
