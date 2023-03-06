package main

import (
	"fmt"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/yihsuanhung/go-line-bot/internal/handler"
)

func main() {
	// set up db
	// err := db.Init()
	// if err != nil {
	// 	panic(err)
	// }

	// config line bot
	// var err error
	// bot, err = linebot.New(channelSecret, channelAccessToken)
	// if err != nil {
	// 	log.Fatalf("Failed to create LineBot client: %v", err)
	// }

	// set up routes
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/message", handler.CreateUserMessage)
		v1.POST("/webhook", handler.Webhook)
	}
	// r.POST("/message", handler.CreateUserMessage)
	// r.POST("/webhook", handler.Webhook)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start gin: %v", err)
	}

	// config := server.DefaultConfig()
	// instance := config.Build()
	// v1 := instance.Group("v1")
	// v1.GET("/hello", handler.Hello)
	// v1.POST("/webhook", handler.Webhook)
	// v1.OPTIONS("/webhook", handler.Preflight) // TODO: middleware
	// if err := instance.Serve(); err != nil {
	// 	panic(err)
	// }

}
