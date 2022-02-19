package response

import (
	errNo "fagin/app/errno"
	"fagin/pkg/errno"
	"fagin/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

const DefCap = 20

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

type Model struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func JSONSuccess(ctx *gin.Context, result interface{}) *Model {
	return JSON(ctx, errNo.OK, result, nil, http.StatusOK)
}

func JSONErr(ctx *gin.Context, err error, errors interface{}) *Model {
	return JSON(ctx, err, nil, errors, http.StatusOK)
}

func JSONWithStatus(ctx *gin.Context, statusCode int, err error, result, errors interface{}) *Model {
	return JSON(ctx, err, result, errors, statusCode)
}

func JSON(ctx *gin.Context, err error, result, errors interface{}, statusCode int) *Model {
	code, msg := errno.Decode(err)
	res := Model{
		Code:    code,
		Message: msg,
		Result:  result,
		Errors:  errors,
	}
	ctx.JSON(statusCode, res)

	return &res
}

func (res *Model) Log(model string, args ...interface{}) {
	args = append([]interface{}{"响应信息: ", res.Message}, args...)
	go logger.Channel(model).Info(args...)
}
