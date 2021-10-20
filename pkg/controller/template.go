package controller

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateControllerTemplate(path, name string) {
	filePath := config.App().AppPath + "/controllers/" + path + ".go"
	sl := strings.Split(filePath, "/")
	dirPath := strings.Join(sl[:len(sl)-1], "/")
	packageName := sl[len(sl)-2]
	name = utils.Camel(name)
	structName := strings.ToLower(string(name[0])) + name[1:]

	//os.Stat获取文件信息
	if _, err := os.Stat(filePath); err == nil {
		panic("文件已存在")
	}

	// 创建文件夹 可以多层
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	const temp = `package %[1]s

import (
	"fagin/app/errno"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type %[2]s struct{
	BaseController
}

var %[3]s %[2]s

// Index 列表
func (ctr *%[2]s) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	params := gin.H{
		"orderBy": "id asc",
	}
	columns := []string{"id"}
	with := gin.H{}
	result, err := service.S.Index(params, columns, with, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.R(result...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"data": data,
		"paginator":  paginator,
	})
	return
}

// Show 展示
func (ctr *%[2]s) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id"}
	data, err := service.S.Show(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"id": data.ID,
	})
	return
}

// Store 创建
func (ctr *%[2]s) Store(ctx *gin.Context) {
	var r r.R
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	m := model{}

	err := service.S.Create(&m)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Update 更新
func (ctr *%[2]s) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r r.Create
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	data := map[string]interface{}{}
	err = service.S.Update(id, data)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Del 删除
func (ctr *%[2]s) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.S.Delete(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Deletes 批量删除
func (ctr *%[2]s) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint ` + "`form:\"ids\" json:\"ids\" binding:\"required\"`" + `
	}
	var ids R
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.S.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	c.ResponseJsonOK(ctx, nil)
	return
}
`
	content := fmt.Sprintf(temp, packageName, structName, name, config.App().Name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("controller create run successfully")
}
