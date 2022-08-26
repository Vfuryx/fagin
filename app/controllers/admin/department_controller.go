package admin

import (
	"fagin/app/errno"
	"fagin/app/models/admin_department"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type departmentController struct {
	BaseController
}

// DepartmentController 部门控制器
var DepartmentController departmentController

// Index 列表
func (ctr *departmentController) Index(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.AdminDepartmentList](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	params := gin.H{}
	if r.Name != "" {
		params["like_name"] = "%" + r.Name + "%"
	}

	if r.Status != nil {
		params["status"] = *r.Status
	}

	columns := []string{"*"}
	with := gin.H{}

	result, err := services.AdminDepartment.Index(params, columns, with)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.CtxListErr, nil)
		return
	}

	data := response.NewAdminDepartment(result...).Collection()

	ctr.ResponseJSONOK(ctx, data)
}

// Show 展示
func (ctr *departmentController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id"}

	data, err := services.AdminDepartment.Show(id, columns)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"id":        data.ID,
		"parent_id": data.ParentID,
		"name":      data.Name,
		"remark":    data.Remark,
		"sort":      data.Sort,
		"status":    data.Status,
	})
}

// Store 创建
func (ctr *departmentController) Store(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.CreateAdminDepartment](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	m := admin_department.AdminDepartment{
		ParentID: r.ParentID,
		Name:     r.Name,
		Remark:   r.Remark,
		Sort:     r.Sort,
		Status:   *r.Status,
	}

	err := services.AdminDepartment.Create(&m)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Update 更新
func (ctr *departmentController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	r, msg := request.Validation[adminRequest.UpdateAdminDepartment](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	data := map[string]interface{}{
		"parent_id": r.ParentID,
		"name":      r.Name,
		"remark":    r.Remark,
		"sort":      r.Sort,
		"status":    r.Status,
	}

	err = services.AdminDepartment.Update(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Del 删除
func (ctr *departmentController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.AdminDepartment.Delete(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Deletes 批量删除
func (ctr *departmentController) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint `form:"ids" json:"ids" binding:"required"`
	}

	var ids R

	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.AdminDepartment.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}
