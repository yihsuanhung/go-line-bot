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
		v1.GET("/message", handler.GetAllMessages)
		v1.GET("/message/:id", handler.GetMessageByID)
		v1.POST("/message", handler.CreateMessage)
		v1.PUT("/message/:id", handler.UpdateMessage)
		v1.DELETE("/message/:id", handler.DeleteMessage)

		// commit #5 Webhook
		v1.POST("/webhook", handler.Webhook)

		// commit #6 Send message back to line
		v1.POST("/message-delivery", handler.SendMessage)

		// commit #7 Query message list of the user
		// v1.GET("/user-message/:userID", handler.GetMessagesByUserID)

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
