package controller

import (
	"fagin/app"
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
	name = app.Camel(name)
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
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type %[2]s struct{}

var %[3]s %[2]s

// 列表
func (%[2]s) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := gin.H{
		"sort": "id asc",
	}
	columns := []string{"id"}
	with := gin.H{"Model": nil}
	result, err := service.S.Index(params, columns, with, &paginator)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrList, nil)
		return
	}

	data := response.R(result...).Collection()

	app.JsonResponse(ctx, errno.OK, gin.H{
		"data": data,
		"paginator":  paginator,
	})
	return
}

// 展示
func (%[2]s) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	columns := []string{"id"}
	data, err := service.S.Show(id, columns)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.Err, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{
		"id":     data.ID,
	})
	return
}

// 创建
func (%[2]s) Store(ctx *gin.Context) {
	var r r.R
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	c := model{}

	err := service.S.Create(&c)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.Err, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 更新
func (%[2]s) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r r.Create
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}
	data := map[string]interface{}{}
	err = service.S.Update(id, data)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrUpdate, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 删除
func (%[2]s) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.S.Delete(id)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrDelete, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 批量删除
func (%[2]s) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint `+"`form:\"ids\" json:\"ids\" binding:\"required\"`"+`
	}
	var ids R
	err := ctx.ShouldBind(&ids)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.S.Deletes(ids.IDs)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrDelete, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
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