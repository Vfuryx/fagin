package services

import (
	aol "fagin/app/models/admin_operation_log"
	"fagin/pkg/db"
	"fagin/pkg/errorw"

	"github.com/gin-gonic/gin"
)

type adminOperationLog struct{}

var AdminOperationLog adminOperationLog

func (*adminOperationLog) List(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]*aol.AdminOperationLog, error) {
	var operationLogs []*aol.AdminOperationLog
	err := aol.NewDao().Query(params, columns, with).Paginate(&operationLogs, p)

	return operationLogs, errorw.UP(err)
}

func (*adminOperationLog) ShowLog(id uint, columns []string) (*aol.AdminOperationLog, error) {
	log := aol.New()
	err := log.Dao().FindByID(id, columns)

	return log, errorw.UP(err)
}
