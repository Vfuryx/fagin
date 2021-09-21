package response

import (
	errNo "fagin/app/errno"
	"fagin/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IResponse interface {
	Serialize() []map[string]interface{}
	Handle() []map[string]interface{}
	Item() map[string]interface{}
	Collection() []map[string]interface{}
}

type Collect struct {
	IResponse
}

func (c *Collect) SetCollect(res IResponse) {
	c.IResponse = res
}

func (c Collect) Handle() []map[string]interface{} {
	return c.Serialize()
}

func (c Collect) Item() map[string]interface{} {
	return c.Handle()[0]
}

func (c Collect) Collection() []map[string]interface{} {
	return c.Handle()
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func JsonOK(ctx *gin.Context, data interface{}) {
	Json(ctx, errNo.OK, data, nil, http.StatusOK)
}

func JsonErr(ctx *gin.Context, err error, errors interface{}) {
	Json(ctx, err, nil, errors, http.StatusOK)
}

func JsonWithStatus(ctx *gin.Context, statusCode int, err error, data interface{}, errors interface{}) {
	Json(ctx, err, data, errors, statusCode)
}

func Json(ctx *gin.Context, err error, data interface{}, errors interface{}, statusCode int) {
	code, msg := errno.Decode(err)
	ctx.JSON(statusCode, Response{
		Code:    code,
		Message: msg,
		Data:    data,
		Errors:  errors,
	})
}
