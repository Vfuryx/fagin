package response

import (
	errNo "fagin/app/errno"
	"fagin/pkg/errno"
	"fagin/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response interface {
	Serialize() []map[string]interface{}
	Handle() []map[string]interface{}
	Item() map[string]interface{}
	Collection() []map[string]interface{}
}

type Collect struct {
	Response
}

func (c *Collect) SetCollect(res Response) {
	c.Response = res
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

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func JsonOK(ctx *gin.Context, result interface{}) *response {
	return Json(ctx, errNo.OK, result, nil, http.StatusOK)
}

func JsonErr(ctx *gin.Context, err error, errors interface{}) *response {
	return Json(ctx, err, nil, errors, http.StatusOK)
}

func JsonWithStatus(ctx *gin.Context, statusCode int, err error, result interface{}, errors interface{}) *response {
	return Json(ctx, err, result, errors, statusCode)
}

func Json(ctx *gin.Context, err error, result interface{}, errors interface{}, statusCode int) *response {
	code, msg := errno.Decode(err)
	res := response{
		Code:    code,
		Message: msg,
		Result:  result,
		Errors:  errors,
	}
	ctx.JSON(statusCode, res)
	return &res
}

func (res *response) Log(model string, args ...interface{}) {
	args = append([]interface{}{"响应信息", res.Message}, args...)
	go logger.Channel(model).Info(args...)
}
