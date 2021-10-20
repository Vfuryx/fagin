package service

import (
	all "fagin/app/models/admin_login_log"
	"fagin/pkg/db"
	"fagin/pkg/errorw"
	"fagin/utils"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type adminLoginLog struct{}

// AdminLoginLog 后台登录日志服务
var AdminLoginLog adminLoginLog

func (*adminLoginLog) List(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]all.AdminLoginLog, error) {
	var loginLogs []all.AdminLoginLog
	err := all.NewDao().Query(params, columns, with).Paginate(&loginLogs, p)
	return loginLogs, errorw.UP(err)
}

func (*adminLoginLog) ShowLog(id uint, columns []string) (*all.AdminLoginLog, error) {
	log := all.New()
	err := log.Dao().FindById(id, columns)
	return log, errorw.UP(err)
}

// LogStore 记录登录日志
func (*adminLoginLog) LogStore(username, ip, userAgent string) error {
	ua := user_agent.New(userAgent)
	browser, _ := ua.Browser()
	log := &all.AdminLoginLog{
		LoginName: username,
		IP:        ip,
		Location:  utils.GetLocation(ip),
		Browser:   browser,
		OS:        ua.OS(),
		Status:    1,
	}
	err := all.NewDao().Create(log)
	return errorw.UP(err)
}
