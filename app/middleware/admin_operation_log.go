package middleware

import (
	"bytes"
	"encoding/json"
	"fagin/app"
	"fagin/app/models/admin_user"
	"fagin/app/mq"
	"fagin/config"
	"fagin/pkg/auths"
	"fagin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"io"
	"runtime/debug"
	"strings"
	"time"
)

var excludeMethods = map[string]struct{}{
	"GET":     {},
	"OPTIONS": {},
}

// AdminOperationLogToDB 日志记录到文件
func AdminOperationLogToDB() func(*gin.Context) {

	return func(ctx *gin.Context) {
		// 请求方式
		reqMethod := ctx.Request.Method

		if _, ok := excludeMethods[reqMethod]; ok {
			ctx.Next()
		} else if strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/json") { // 只记录 json 参数
			// 获取body
			data, _ := ctx.GetRawData()
			//恢复body
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 关键点
			// 开始时间
			startTime := time.Now()
			// 处理请求
			ctx.Next()

			adminOperationLogToDB(ctx, reqMethod, startTime, data)
		} else {
			ctx.Next()
		}

	}
}

func adminOperationLogToDB(ctx *gin.Context, reqMethod string, startTime time.Time, body []byte) {
	// 执行时间
	latencyTime := time.Now().Sub(startTime)

	// 请求IP
	clientIP := ctx.ClientIP()
	// 获取用户ID
	uid := auths.GetAdminID(ctx)
	//// 获取类型
	//logType := ctx.GetInt(admin_operation_log.TypeKey)
	//// 获取内容
	//content := ctx.GetString(admin_operation_log.ContentKey)

	if config.AMQP.Open {
		adminLog := mq.AdminLog{
			LatencyTime: latencyTime,
			Body:        body,
			Method:      reqMethod,
			ClientIP:    clientIP,
			UserName:    ctx.GetString(admin_user.AdminUserNameKey),
			UserAgent:   ctx.Request.UserAgent(),
			URI:         ctx.Request.RequestURI,
			Path:        ctx.FullPath(),
			AdminID:     uint(uid),
			//Content:     content,
			//LogType:     uint8(logType),
		}
		go func(adminLog mq.AdminLog) {
			b, err := json.Marshal(adminLog)
			if err != nil {
				go app.Log(logger.AdminModel).Println(err, string(debug.Stack()))
				return
			}

			err = mq.AdminLogMQ.Publish(amqp.Publishing{
				ContentType: "application/json",
				Body:        b,
			})
			if err != nil {
				go app.Log(logger.AdminModel).Println(err, string(debug.Stack()))
				return
			}
		}(adminLog)
	} else {
		go mq.LogStore(
			body,
			latencyTime,
			uint(uid),
			reqMethod,
			clientIP,
			ctx.GetString(admin_user.AdminUserNameKey),
			ctx.Request.UserAgent(),
			ctx.Request.RequestURI,
			ctx.FullPath(),
		)
	}
}

//// AdminOperationLogSetType 日志记录到文件
//func AdminOperationLogSetType(logType int) func(*gin.Context) {
//	return func(ctx *gin.Context) {
//		if _, ok := excludeMethods[ctx.Request.Method]; ok {
//			ctx.Next()
//		} else {
//			ctx.Set(admin_operation_log.TypeKey, logType)
//		}
//	}
//}
