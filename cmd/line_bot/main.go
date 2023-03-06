package main

import (
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/yihsuanhung/go-line-bot/pkg/handler"
)

func main() {
	// Set up routes
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		// User message CRUD endpoints
		v1.GET("/user-message", handler.GetAllUserMessages)
		v1.GET("/user-message/:id", handler.GetUserMessageByID)
		v1.POST("/user-message", handler.CreateUserMessage)
		v1.PUT("/user-message/:id", handler.UpdateUserMessage)
		v1.DELETE("/user-message/:id", handler.DeleteUserMessage)

		// commit #5 Webhook
		v1.POST("/webhook", handler.Webhook)

		// commit #6 Send message back to line
		v1.POST("/message", handler.SendMessage)

		// commit #7 Query message list of the user
		//TODO
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	if err := r.Run(addr); err != nil {
		panic(fmt.Sprintf("Failed to start gin: %v", err))
	}
}
