package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yihsuanhung/go-line-bot/internal/model"
)

var bot *linebot.Client

func Webhook(c *gin.Context) {

	if bot != nil {
		return
	}

	var err error
	bot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			fmt.Println("Invalid Signature", err)
			c.AbortWithStatus(400)
		} else {
			fmt.Println("Error", err)
			c.AbortWithStatus(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// *Handle only on text message
			case *linebot.TextMessage:
				userID := event.Source.UserID
				profile, err := bot.GetProfile(userID).Do()

				if err != nil {
					fmt.Println("Get profile err:", err)
				}

				_, err = bot.GetMessageQuota().Do()
				if err != nil {
					fmt.Println("Quota err:", err)
				}

				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("OK! 收到訊息: %s \n來自: %s", message.Text, profile.DisplayName))).Do()

				userMessage := &model.UserMessage{
					MessageId: message.ID,
					Message:   message.Text,
					UserId:    profile.UserID,
					UserName:  profile.DisplayName,
				}

				model.CreateUserMessage(userMessage)

			// *Handle only on Sticker message
			case *linebot.StickerMessage:

				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OK! 收到貼圖")).Do()

				// var kw string
				// for _, k := range message.Keywords {
				// 	kw = kw + "," + k
				// }
				// outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerID, message.PackageID, kw, message.Text)
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
				// 	fmt.Print(err)
				// }
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}