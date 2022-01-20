# Lets create a simple web server using Gin.

package main
import "github.com/gin-gonic/gin"

func main() {

	// Creates a gin router with default middleware
	router  := gin.Default()

	// A handler for GET request on /example
	router.GET("/example", func(c *gin.Context) {
		
		c.JSON(200, gin.H{
			"message": "example",
		}) // gin.H is a shortcut for map[string]interface{}
	})
	router.Run() // listen and serve on port 8080
}

gin.Defalut() creates a Gin router with default middleware: logger & recover middleware

We can make handler using router.GET(path, handle) where path is the relative path, and handle is the handler function that takes *gin.Context as an argument. The handler function serves a JSON response with a status of 200.

And finally we can start the router using router.Run() //By defalut it listens on port 8080


# H is shortcut for map[string]interface{}
type H map[string]interface{}

This is to say,
c.JSON(http.StatusOK, gin.H {"status":"login successful"})

Equivalent to,
c.JSON(http.StatusOK, map [string] interface {} {"status": "login successful"})

Ex: If you need to nest JSON, then nesting gin.H will do the job

c.JSON(http.StatusOK, gin.H{
    "status": gin.H{
        "code": http.StatusOK,
        "status": "Login succesful"
    },
    "message": message
})


# IndentedJSON - Formats the JSON response
c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})