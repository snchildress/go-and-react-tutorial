package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup router group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "APIs working as expected!",
			})
		})
	}

	// Joke API routes
	api.GET("/jokes", RetrieveJokesHandler)
	api.POST("/jokes/likes/:jokeID", LikeJokeHandler)

	// Start and run the server
	router.Run(":3000")
}

// RetrieveJokesHandler retrieves a list of available jokes
func RetrieveJokesHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Jokes retrieving handler not yet implemented",
	})
}

// LikeJokeHandler likes a given joke
func LikeJokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Joke liking handler not yet implemented",
	})
}
