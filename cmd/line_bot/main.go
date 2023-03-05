package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yihsuanhung/go-line-bot/internal/controller"
	"github.com/yihsuanhung/go-line-bot/internal/model"
)

// var bot *linebot.Client

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
	r.POST("/message", controller.CreateMessage)
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

	// channelSecret := "8723770c0c4e010ae92251a917553bf2"

	// defer c.Request.Body.Close()
	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	http.Error(w, "Error reading request body", http.StatusInternalServerError)
	// 	return
	// }
	// decoded, err := base64.StdEncoding.DecodeString(c.Request.Header.Get("x-line-signature"))
	// if err != nil {
	// 	// ...
	// }
	// hash := hmac.New(sha256.New, []byte(channelSecret))
	// hash.Write(body)

	// fmt.Println(hmac.Equal(hash.Sum(nil), decoded))

	// channelSecret := ""
	// channelAccessToken := ""

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Println(err)
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
			// Handle only on text message
			case *linebot.TextMessage:
				userID := event.Source.UserID
				profile, err := bot.GetProfile(userID).Do()

				if err != nil {
					log.Println("Get profile err:", err)
				}

				_, err = bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}

				// message.ID: Msg unique ID
				// message.Text: Msg text
				// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("msg ID:"+message.ID+":"+"Get:"+message.Text+" , \n OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
				// 	log.Print(err)
				// }

				// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				// defer cancel()

				// fmt.Println("空db", db.Collection == nil)
				// db.Collection.InsertOne(ctx, bson.D{{Key: "name", Value: "pi"}, {Key: "value", Value: 11111}})

				// model.CreateMessage()

				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("OK! 收到訊息: %s \n來自: %s", message.Text, profile.DisplayName))).Do()

				userMessage := &model.UserMessage{
					MessageId: message.ID,
					Message:   message.Text,
					UserId:    profile.UserID,
					UserName:  profile.DisplayName,
				}

				model.CreateUserMessage(userMessage)

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
