package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		for l, i := range c.Req.Header {
			fmt.Fprintf(c.Writer, "header[%q] = %q\n", l, i)
		}
	})
	r.Run(":9999")
}
