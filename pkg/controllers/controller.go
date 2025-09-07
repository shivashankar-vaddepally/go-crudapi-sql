package controllers

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vaddepally-shiva-shankar/go-crudapi-sql/pkg/models"
	"github.com/vaddepally-shiva-shankar/go-crudapi-sql/pkg/utils"
)

var NewBook models.Book

func GetBooks(c *gin.Context) {
	fmt.Println("Entering into GetBooks function")
	booksDetails := NewBook.GetAllBooks()
	res, err := utils.MarshalFunc(booksDetails)
	if err != nil {
		fmt.Println("error while marshal")
		c.JSON(500, gin.H{"error": "failed to marshal details"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
	return
}

func GetBookbyId(c *gin.Context) {
	id := c.Params.ByName("ID")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")

	}
	bookDetails := NewBook.GetBookbyId(ID)
	res, err := utils.MarshalFunc(bookDetails)
	if err != nil {
		fmt.Println("error while marshal")
		c.JSON(500, gin.H{"error": "failed to marshal details"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
	return
}

func CreateBook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error while reading body")
		c.JSON(500, gin.H{"error": "failed to read request body"})
		return
	}
	CreateBook, err := utils.UnmarshalFunc([]byte(body))
	if err != nil {
		fmt.Println("error while unmarshal")
		c.JSON(500, gin.H{"error": "failed to unmarshal details"})
		return
	}
	b := CreateBook.CreateBook()
	res, err := utils.MarshalFunc(b)
	if err != nil {
		fmt.Println("error while marshal")
		c.JSON(500, gin.H{"error": "failed to marshal details"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
	return
}

func DeleteBook(c *gin.Context) {
	id := c.Params.ByName("ID")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		c.JSON(500, gin.H{"error": "failed to parse id"})
		return
	}
	bookDetails := NewBook.DeleteBook(ID)
	res, err := utils.MarshalFunc(bookDetails)
	if err != nil {
		fmt.Println("error while marshalling")
		c.JSON(500, gin.H{"error": "failed to marshal details"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
	return
}

// func UpdateBook(c *gin.Context) {
// 	id := c.Params.ByName("ID")
// 	ID, err := strconv.ParseInt(id, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")

// 	}
// 	bookDetails := NewBook.DeleteBook(ID)

// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		fmt.Println("error while reading body")
// 		c.JSON(500, gin.H{"error": "failed to read request body"})
// 		return
// 	}
// 	CreateBook, err := utils.UnmarshalFunc([]byte(body))
// 	if err != nil {
// 		fmt.Println("error while unmarshal")
// 		c.JSON(500, gin.H{"error": "failed to unmarshal details"})
// 		return
// 	}
// 	b := CreateBook.CreateBook()
// 	res, err := utils.MarshalFunc(b)
// 	if err != nil {
// 		fmt.Println("error while marshal")
// 		c.JSON(500, gin.H{"error": "failed to marshal details"})
// 		return
// 	}
// 	c.Header("Content-Type", "application/json")
// 	c.Writer.WriteHeader(200)
// 	c.Writer.Write(res)
// 	return

// }
