package main

import (
	"fmt"
	lgr "github.com/XFroggyX/go-logger"
	config "github.com/XFroggyX/my-little-news-magazine/internal/config"
	"github.com/XFroggyX/my-little-news-magazine/internal/user"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	router := gin.New()
	lgr.Init("logs", "news-magazine.log")
	logger := lgr.GetLogger()

	cfg := config.GetConfig()

	logger.Info("create handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	run(router, cfg)
}

func run(router http.Handler, cfg *config.Config) {
	logger := lgr.GetLogger()
	log.Println("Server run")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		sockPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", sockPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", sockPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.Serve(listener)
	if err != nil {
		logger.Fatal(err)
	}
}
