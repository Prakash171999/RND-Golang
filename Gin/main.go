package main

import ("fmt"
	"github.com/gin-gonic/gic"
)

func main(){
	fmt.Println("Hello world")

	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	r.Run()
}