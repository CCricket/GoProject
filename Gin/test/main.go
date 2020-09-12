package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"小王子",
		})
	})

	r.POST("/post", func(c *gin.Context){
		c.JSON(200, gin.H{
			"Message":"Ivana",
		})
	})
	r.PUT("/PUT", func(c *gin.Context){
		c.JSON(200, gin.H{
			"PUT":8899,
		})
	})
	r.DELETE("/delete", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"delete",
		})
	})
	r.PATCH("/patch", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"PATCH",
		})
	})
	r.HEAD("/head", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"head",
		})
	})
	/*r.Handle("DELETE", "/delete", func(c gin.Context){
		c.String(200, "delete")
	})*/
	r.OPTIONS("/options", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"options",
		})
	})

	r.Any("/any", func(c *gin.Context){
		c.String(200,"any")
	})
	r.Run()
}
