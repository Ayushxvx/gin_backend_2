package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var posts = []Post{
	{ID: "1", Title: "Kenny and Darla", Content: "Wheaton Valley High, class of '78", CreatedAt: time.Date(2001, time.May, 1, 22, 0, 0, 0, time.UTC)},
}

var pl = fmt.Println

func main() {
	pl("Server is starting...")
	router := gin.Default()
	router.Use(cors.Default())

	gin.SetMode(gin.ReleaseMode)
	router.GET("/", showPosts)

	router.POST("/addpost", addPosts)

	router.Run("0.0.0.0")

}

func showPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

func addPosts(c *gin.Context) {
	var newPost Post

	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost.ID = strconv.Itoa((len(posts) + 1))
	newPost.CreatedAt = time.Now()
	posts = append(posts, newPost)
	c.JSON(http.StatusCreated, newPost)
}
