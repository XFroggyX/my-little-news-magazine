package main

import (
	lgr "github.com/XFroggyX/go-logger"
	"github.com/XFroggyX/my-little-news-magazine/internal/user"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	router := gin.New()
	lgr.Init("logs", "news-magazine.log")
	logger := lgr.GetLogger()
	logger.Info("create handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	run(router)
}

func run(router *gin.Engine) {
	logger := lgr.GetLogger()
	log.Println("Server run")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Error("listener error")
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = server.Serve(listener)
	if err != nil {
		logger.Fatal(err)
	}
}
