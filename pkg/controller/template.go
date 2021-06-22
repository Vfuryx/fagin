package controller

import (
	"fagin/app/utils"
	"fagin/config"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateControllerTemplate(path, name string) {
	filePath := config.App.AppPath + "/controllers/" + path + ".go"
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
func (c *%[2]s) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := gin.H{
		"orderBy": "id asc",
	}
	columns := []string{"id"}
	with := gin.H{"Model": nil}
	result, err := service.S.Index(params, columns, with, &paginator)
	if err != nil {
		c.ResponseJsonErr(ctx, errno.Serve.ListErr, nil)
		return
	}

	data := response.R(result...).Collection()

	c.ResponseJsonOK(ctx, gin.H{
		"data": data,
		"paginator":  paginator,
	})
	return
}

// Show 展示
func (c *%[2]s) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		c.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	columns := []string{"id"}
	data, err := service.S.Show(id, columns)
	if err != nil {
		c.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}

	c.ResponseJsonOK(ctx, gin.H{
		"id": data.ID,
	})
	return
}

// Store 创建
func (c *%[2]s) Store(ctx *gin.Context) {
	var r r.R
	if data, ok := r.Validate(ctx); !ok {
		c.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	m := model{}

	err := service.S.Create(&m)
	if err != nil {
		c.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	c.ResponseJsonOK(ctx, nil)
	return
}

// Update 更新
func (c *%[2]s) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		c.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r r.Create
	if data, ok := r.Validate(ctx); !ok {
		c.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}
	data := map[string]interface{}{}
	err = service.S.Update(id, data)
	if err != nil {
		c.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	c.ResponseJsonOK(ctx, nil)
	return
}

// Del 删除
func (c *%[2]s) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		c.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.S.Delete(id)
	if err != nil {
		c.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	c.ResponseJsonOK(ctx, nil)
	return
}

// Deletes 批量删除
func (c *%[2]s) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint `+"`form:\"ids\" json:\"ids\" binding:\"required\"`"+`
	}
	var ids R
	err := ctx.ShouldBind(&ids)
	if err != nil {
		c.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.S.Deletes(ids.IDs)
	if err != nil {
		c.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	c.ResponseJsonOK(ctx, nil)
	return
}
`
	content := fmt.Sprintf(temp, packageName, structName, name, config.App.Name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("controller create run successfully")
}