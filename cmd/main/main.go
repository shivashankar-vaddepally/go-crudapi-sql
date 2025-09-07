package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vaddepally-shiva-shankar/go-crudapi-sql/pkg/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterBookstoreRouter(r)
	log.Fatal(r.Run(":8080"))
	fmt.Println("listening on port 8080")
}
