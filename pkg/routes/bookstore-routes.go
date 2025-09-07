package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vaddepally-shiva-shankar/go-crudapi-sql/pkg/controllers"
)

func RegisterBookstoreRouter(router *gin.Engine) {
	fmt.Printf("Entering into RegisterBookstoreRouter function\n")
	log.Printf("Entering into RegisterBookstoreRouter function\n")
	router.GET("/book", controllers.GetBooks)
	router.GET("book/:id", controllers.GetBookbyId)
	router.POST("/book", controllers.CreateBook)
	//router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)
}
