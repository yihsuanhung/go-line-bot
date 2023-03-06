package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yihsuanhung/go-line-bot/internal/model"
)

func GetAllMessages(c *gin.Context) {
	msg, err := model.GetAllMessages()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, msg)
	}
}

func GetMessageByID(c *gin.Context) {
	id := c.Params.ByName("id")
	msg, err := model.GetMessageByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, msg)
	}
}

func CreateMessage(c *gin.Context) {
	var msg model.Message
	if err := c.BindJSON(&msg); err != nil {
		fmt.Println("Bind JSON error", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if err := model.CreateMessage(&msg); err != nil {
			fmt.Println("Create message error", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusCreated, msg)
		}
	}
}

func UpdateMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	msg, err := model.GetMessageByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		if err := c.BindJSON(&msg); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			if err := model.UpdateMessage(msg); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			} else {
				c.JSON(http.StatusOK, msg)
			}
		}
	}
}

func DeleteMessage(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := model.DeleteMessage(id); err != nil {
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

func GetMessagesByUserId(c *gin.Context) {
	userId := c.Params.ByName("userId")
	msg, err := model.GetMessagesByUserId(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, msg)
	}

}
