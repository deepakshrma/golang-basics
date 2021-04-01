// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deepakshrma/todo_app/repository"
	"github.com/gin-gonic/gin"
)

var (
	r = repository.NewRepository()
)

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": r.ListTodo()})
}

func addTodo(c *gin.Context) {
	var input repository.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r.AddTodo(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}
func main() {
	// Get the user provided port from ENV, if port is not defined user default port as 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/todos", getTodos)
		v1.POST("/todos", addTodo)
	}
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
