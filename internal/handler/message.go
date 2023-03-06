package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yihsuanhung/go-line-bot/internal/model"
)

func GetAllUserMessages(c *gin.Context) {
	userMessages, err := model.GetAllUserMessages()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, userMessages)
	}
}

func GetUserMessageByID(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := model.GetUserMessageByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUserMessage(c *gin.Context) {
	var message model.UserMessage
	if err := c.BindJSON(&message); err != nil {
		fmt.Println("Bind JSON error", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if err := model.CreateUserMessage(&message); err != nil {
			fmt.Println("Create message error", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusCreated, message)
		}
	}
}

func UpdateUserMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	userMessage, err := model.GetUserMessageByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		if err := c.BindJSON(&userMessage); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			if err := model.UpdateUserMessage(userMessage); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			} else {
				c.JSON(http.StatusOK, userMessage)
			}
		}
	}
}

func DeleteUserMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := model.DeleteUserMessage(id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

func SendMessage(c *gin.Context) {
	var sendingMessage model.SendingMessage
	if err := c.BindJSON(&sendingMessage); err != nil {
		fmt.Println("Bind JSON error", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if err := model.SendMEssage(&sendingMessage); err != nil {
			fmt.Println("Sending message error", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
	}
}
