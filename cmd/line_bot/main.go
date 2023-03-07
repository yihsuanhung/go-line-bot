package main

import (
	"fmt"
	"github.com/yihsuanhung/go-line-bot/internal/cmd"
	"github.com/yihsuanhung/go-line-bot/pkg/bot"

	"github.com/gin-gonic/gin"
	"github.com/yihsuanhung/go-line-bot/pkg/conf"
	"github.com/yihsuanhung/go-line-bot/pkg/handler"
	"github.com/yihsuanhung/go-line-bot/pkg/settings"
)

func main() {
	cmd.Execute()
	cfg := conf.NewConfig(
		conf.WithPath("./chat/app"), // ./chat/app ../../chat/app
		conf.WithType("yaml"),
	)
	var botInfo settings.LineBotInfo
	cfg.Unmarshal(&botInfo)
	bot.InitLineBot(&botInfo)
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
		v1.GET("/user-message/:userId", handler.GetMessagesByUserId)
	}
	addr := fmt.Sprintf(":%s", cmd.Port)
	if err := r.Run(addr); err != nil {
		panic(fmt.Sprintf("Failed to start gin: %v", err))
	}
}
