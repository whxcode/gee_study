package main

import (
	"fmt"

	"gee/src/gee"
)

func main() {
	h := gee.New()

	h.GET("/json", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"message": "hello json",
		})
	})

	h.GET("/", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.Writer, "%s: %s\n", k, v)
		}
	})

	h.Run(":8080")
}
