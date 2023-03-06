package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yihsuanhung/go-line-bot/internal/model"
	"github.com/yihsuanhung/go-line-bot/pkg/bot"
)

func Webhook(c *gin.Context) {
	events, err := bot.LineBot.ParseRequest(c.Request)
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
				userId := event.Source.UserID
				profile, err := bot.LineBot.GetProfile(userId).Do()

				if err != nil {
					fmt.Println("Get profile err:", err)
				}

				_, err = bot.LineBot.GetMessageQuota().Do()
				if err != nil {
					fmt.Println("Quota err:", err)
				}

				bot.LineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("OK! 收到訊息: %s \n來自: %s", message.Text, profile.DisplayName))).Do()

				userMessage := &model.Message{
					MessageId: message.ID,
					Message:   message.Text,
					UserId:    profile.UserID,
					UserName:  profile.DisplayName,
				}

				model.CreateMessage(userMessage)

			// *Handle only on Sticker message
			case *linebot.StickerMessage:
				bot.LineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OK! 收到貼圖")).Do()
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
