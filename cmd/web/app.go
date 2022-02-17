package main

import (
	"github.com/XFroggyX/my-little-news-magazine/internal/user"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	handler := user.NewHandler()
	handler.Register(router)

	run(router)
}

func run(router *gin.Engine) {
	log.Println("Server run")
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
