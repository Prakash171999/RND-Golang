package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []book{
	{ID: "1", Title: "Harry Potter: Goblet of fire", Author: "J.K Rowling"},
	{ID: "2", Title: "Rich Dad Poor Dad", Author: "Robert Kiyosaki"},
	{ID: "3", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
}

func main() {
	r := gin.Default()
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookById)
	r.POST("/books/", postBooks)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/delete/:id", deleteBook)
	r.Run()
}

//GET BOOKS
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}

//GET BOOKS BY id
func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, b := range Books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

//POST Books BY
func postBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	Books = append(Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

//UPDATE BOOKS
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var newBook book

	//get body in newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	//iterate the slice and update the book
	for index, b := range Books {
		if b.ID == id {
			if newBook.Author != b.Author && newBook.Author != "" {
				Books[index].Author = newBook.Author
			}
			if newBook.Title != b.Title && newBook.Title != "" {
				Books[index].Title = newBook.Title
			}
		}
	}
	//Sending response
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

//DELETE BOOK
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	//iterate and delete the element
	for index, book := range Books {
		if book.ID == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}

	for _, book := range Books {
		fmt.Println(book.ID, book.Title, book.Author)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
