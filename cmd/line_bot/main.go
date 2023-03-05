package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yihsuanhung/go-line-bot/internal/db"
)

var bot *linebot.Client

func main() {
	channelSecret := "8723770c0c4e010ae92251a917553bf2"
	channelAccessToken := "2jWKtZiNoEKnvmvL8B4wq0AqXLH0xPZAyBIns+wn51sAYU2B8oYw+0loMNfqvUFWEcnEfezhtVIrCZHqEwQ4/wRNe1Pn4rEvHRpc0SUB53BTQWbKoZJ2RGbqnW35txO2NbjbqNkLkTjQebPtiIXlogdB04t89/1O/w1cDnyilFU="
	// set up db
	db.ConnectDB()

	// config line bot
	var err error
	bot, err = linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Fatalf("Failed to create LineBot client: %v", err)
	}

	// set up routes
	r := gin.Default()
	r.POST("/webhook", webhookHandler)
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

func webhookHandler(c *gin.Context) {

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.AbortWithStatus(400)
		} else {
			c.AbortWithStatus(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// Handle only on text message
			case *linebot.TextMessage:
				// GetMessageQuota: Get how many remain free tier push message quota you still have this month. (maximum 500)
				_, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				// message.ID: Msg unique ID
				// message.Text: Msg text
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("msg ID:"+message.ID+":"+"Get:"+message.Text+" , \n OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
				// 	log.Print(err)
				// }

				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("OK! 收到訊息: %s", message.Text))).Do()

			// Handle only on Sticker message
			case *linebot.StickerMessage:

				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OK! 收到貼圖")).Do()

				// var kw string
				// for _, k := range message.Keywords {
				// 	kw = kw + "," + k
				// }
				// outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerID, message.PackageID, kw, message.Text)
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
				// 	log.Print(err)
				// }
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}
