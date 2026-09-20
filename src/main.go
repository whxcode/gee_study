package main

import (
	"fmt"

	"gee/src/gee"
)

func main() {
	h := gee.New()

	h.GET("/", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.W, "%s: %s\n", k, v)
		}
	})

	h.Run(":8080")
}
