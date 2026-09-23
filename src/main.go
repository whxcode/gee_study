package main

import (
	"fmt"

	"gee/src/gee"
)

func main() {
	h := gee.New()

	h.Use(func(c *gee.Context) {
		fmt.Println("----------------- 开始访问请求 -1-----------------")
		c.Next()
		fmt.Println("----------------- 结束访问请求 -1-----------------")
	})

	{

		user := h.Group("/user")

		user.Use(func(c *gee.Context) {
			fmt.Println("----------------- 开始访问用户请求 0-----------------")
			c.Next()
			fmt.Println("----------------- 结束访问用户请求 0-----------------")
		})

		user.Use(func(c *gee.Context) {
			fmt.Println("----------------- 开始访问用户请求 1-----------------")
			c.Next()
			fmt.Println("----------------- 结束访问用户请求 1-----------------")
		})

		user.GET("/", func(c *gee.Context) {
			fmt.Println("请求用户对接")
			c.String(200, "hellow v1")
		})

		user.GET("/v1", func(c *gee.Context) {
			c.String(200, "user v1\n")
		})

		user.GET("/v1/:name", func(c *gee.Context) {
			c.JSON(200, c.Parmas)
		})
	}

	// ----

	{

		user := h.Group("/group")

		user.GET("/", func(c *gee.Context) {
			c.String(200, "group v1")
		})

		user.GET("/v1", func(c *gee.Context) {
			c.String(200, "group v1\n")
		})

		user.GET("/v1/:name", func(c *gee.Context) {
			c.JSON(200, c.Parmas)
		})
	}

	h.Run(":8080")
}
