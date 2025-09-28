package helper

import "github.com/gin-gonic/gin"

type ResponseWrapper struct {
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func RespondJSON(c *gin.Context, statusCode int, resp ResponseWrapper) {
	c.JSON(statusCode, resp)
	return
}
