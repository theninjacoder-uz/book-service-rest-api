package models

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	IsOk    bool        `json:"isOk"`
	Data    interface{} `json:"data"`
}

func (res *Response) SuccessReponse(c *gin.Context, statusCode int) {
	c.JSON(statusCode, gin.H{
		"message": res.Message,
		"isOk":    res.IsOk,
		"data":    res.Data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"message": err.Error(),
		"isOk":    false,
		"data":    nil,
	})
}
