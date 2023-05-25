package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/hello", hello)
	router.GET("/ilt6", ilt6)
	router.GET("/halo", halo)

	router.Run(":8080")
}

func ping(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "pong"}`),
	)
}

func hello(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "hello ilt 6!"}`),
	)
}

func ilt6(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "this is ilt6!"}`),
	)
}

func halo(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "halo halo!"}`),
	)
}