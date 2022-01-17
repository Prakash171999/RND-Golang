package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
)


func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET - Home page",
	})
}

//Simple Post API
// func PostHomePage(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "POST - Home Page",
// 	})
// }

func PostHomePage(c *gin.Context) {
	//Fetching the data from body
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// c.JSON(200, gin.H{
	// 	"message": "POST - Home Page",
	// })

	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name") // name=elliot
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func PathParameteres(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}


func main(){
	fmt.Println("Hello World")

	r:= gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings) //query?name=elloit&age=24
	r.GET("/path/:name/:age", PathParameteres) // /path/elliot/24/
	r.Run()
}