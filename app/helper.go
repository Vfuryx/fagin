package app

import (
	"fagin/app/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JosnResponse(ctx *gin.Context, err error, data interface{})  {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	code, msg := errno.DecodeErr(err)

	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Message: msg,
		Data: data,
	})
	
}