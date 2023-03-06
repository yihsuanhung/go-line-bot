package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type StatusRequest struct {
// 	ID string `json:"id" binding:"required"`
// }

// type Request interface{}

// deprecated
func Webhook(c *gin.Context) {
	channelSecret := "8723770c0c4e010ae92251a917553bf2"
	// channelAccessToken := "2jWKtZiNoEKnvmvL8B4wq0AqXLH0xPZAyBIns+wn51sAYU2B8oYw+0loMNfqvUFWEcnEfezhtVIrCZHqEwQ4/wRNe1Pn4rEvHRpc0SUB53BTQWbKoZJ2RGbqnW35txO2NbjbqNkLkTjQebPtiIXlogdB04t89/1O/w1cDnyilFU="

	fmt.Println("webhook")

	// * Verify signature
	defer c.Request.Body.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Failed to read request body:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(c.Request.Header.Get("x-line-signature"))
	if err != nil {
		fmt.Println("Failed to decode signature:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hash := hmac.New(sha256.New, []byte(channelSecret))
	hash.Write(body)
	if !hmac.Equal(decoded, hash.Sum(nil)) {
		fmt.Println("Signature does not match")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Println("簽名成功")

	// * Line bot configuration
	// bot, err := linebot.New(channelSecret, channelAccessToken)
	// if err != nil {
	// 	fmt.Println("new bot error")
	// }

	// events, err := bot.ParseRequest(c.Request)
	// if err != nil {
	// 	// Do something when something bad happened.
	// 	fmt.Println("[parse error]", err)
	// }
	// fmt.Println("事件", events)
	// for _, event := range events {
	// 	if event.Type == linebot.EventTypeMessage {
	// 		// Do Something...
	// 		fmt.Println("真事件", events)
	// 	}
	// }

	// * Handle webhook
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// var request Request

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}
