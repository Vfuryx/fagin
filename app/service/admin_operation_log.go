package service

import (
	aol "fagin/app/models/admin_operation_log"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type adminOperationLog struct{}

var AdminOperationLog adminOperationLog

func (adminOperationLog) List(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]aol.AdminOperationLog, error) {
	var operationLogs []aol.AdminOperationLog
	err := aol.Dao().Query(params, columns, with).Paginator(&operationLogs, p)
	return operationLogs, err
}

func (adminOperationLog) ShowLog(id uint, columns []string) (*aol.AdminOperationLog, error) {
	log := aol.New()
	err := log.Dao().FindById(id, columns)
	return log, err
}
