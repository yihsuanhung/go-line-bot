package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yihsuanhung/go-line-bot/internal/model"
)

// func GetAllUsers(c *gin.Context) {
// 	users, err := model.GetAllUsers()
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 	} else {
// 		c.JSON(http.StatusOK, users)
// 	}
// }

// func GetUserByID(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	user, err := model.GetUserByID(id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

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

// func UpdateUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	user, err := model.GetUserByID(id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		if err := c.BindJSON(&user); err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 		} else {
// 			if err := model.UpdateUser(user); err != nil {
// 				c.AbortWithStatus(http.StatusInternalServerError)
// 			} else {
// 				c.JSON(http.StatusOK, user)
// 			}
// 		}
// 	}
// }

// func DeleteUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	if err := model.DeleteUser(id); err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 	} else {
// 		c.Status(http.StatusOK)
// 	}
// }
