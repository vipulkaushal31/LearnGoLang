package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []user{
	{ID: "1", Name: "Vipul Kaushal", Email: "vipulkaushal31@gmail.com", Password: "GoLang@123$"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	fmt.Println("Starting we server")
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")

}
