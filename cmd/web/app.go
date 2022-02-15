package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/:name", indexHandler)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("listener error")
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
