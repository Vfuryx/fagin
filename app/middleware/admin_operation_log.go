package middleware

import (
	"bytes"
	"encoding/json"
	"fagin/app"
	"fagin/app/models/admin_operation_log"
	"fagin/app/service"
	"fagin/pkg/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

type adminOperationLog struct{}

var AdminOperationLog adminOperationLog

// 日志记录到文件
func (adminOperationLog) LoggerToDB() func(*gin.Context) {
	return func(ctx *gin.Context) {
		// 获取body
		data, _ := ctx.GetRawData()
		//恢复body
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点
		var j gin.H
		err := json.Unmarshal(data, &j)
		if err != nil {
			j = gin.H{}
		}

		// 开始时间
		startTime := time.Now()

		// 处理请求
		ctx.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := ctx.Request.Method

		// 请求IP
		clientIP := ctx.ClientIP()

		methods := map[string]struct{}{
			"GET":     {},
			"OPTIONS": {},
		}

		if _, ok := methods[reqMethod]; !ok {
			logger := admin_operation_log.New()
			logger.User = ctx.GetString("user_name")
			logger.Method = reqMethod
			logger.IP = clientIP
			logger.Location = app.GetLocation(clientIP)
			logger.LatencyTime = latencyTime.String()
			logger.UserAgent = ctx.Request.UserAgent()
			logger.Status = 1

			// 请求路由
			logger.Path = ctx.Request.RequestURI
			if logger.Path == "/admin/api/login" {
				logger.Module = "用户登录"
				logger.Operation = "登录"
				logger.User = "-"
			} else if logger.Path == "/admin/api/v1/user/logout" {
				logger.Module = "退出登录"
				logger.Operation = "退出"
			} else {
				if reqMethod == "POST" {
					logger.Operation = "新增"
				} else if reqMethod == "PUT" {
					logger.Operation = "修改"
				} else if reqMethod == "DELETE" {
					logger.Operation = "删除"
				}

				// 模块
				path := ctx.FullPath()
				m, err := service.AdminMenuService.FindByPath(reqMethod, path, []string{"title"})
				if err != nil {
					logger.Module = ""
				}else{
					logger.Module = m.Title
				}
				data, err = json.Marshal(j)
				if err != nil {
					data = []byte("null")
				}
				logger.Input = string(data)
			}

			go func() {
				err := logger.Dao().Create(logger)
				if err != nil {
					log.Log.Errorln(err)
				}
			}()
		}
	}
}
