package admin

import (
	"fagin/app/errno"
	"fagin/app/models/admin_department"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
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
	var r = adminRequest.NewAdminDepartmentList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
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
	result, err := service.AdminDepartment.Index(params, columns, with)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.CtxListErr, nil)
		return
	}

	data := response.AdminDepartment(result...).Collection()

	ctr.ResponseJsonOK(ctx, data)
	return
}

// Show 展示
func (ctr *departmentController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id"}
	data, err := service.AdminDepartment.Show(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"id":        data.ID,
		"parent_id": data.ParentID,
		"name":      data.Name,
		"remark":    data.Remark,
		"sort":      data.Sort,
		"status":    data.Status,
	})
	return
}

// Store 创建
func (ctr *departmentController) Store(ctx *gin.Context) {
	var r = adminRequest.NewCreateAdminDepartment()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	m := admin_department.AdminDepartment{
		ParentID: r.ParentID,
		Name:     r.Name,
		Remark:   r.Remark,
		Sort:     r.Sort,
		Status:   *r.Status,
	}

	err := service.AdminDepartment.Create(&m)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Update 更新
func (ctr *departmentController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = adminRequest.NewUpdateAdminDepartment()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	data := map[string]interface{}{
		"parent_id": r.ParentID,
		"name":      r.Name,
		"remark":    r.Remark,
		"sort":      r.Sort,
		"status":    r.Status,
	}
	err = service.AdminDepartment.Update(id, data)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Del 删除
func (ctr *departmentController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.AdminDepartment.Delete(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Deletes 批量删除
func (ctr *departmentController) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint `form:"ids" json:"ids" binding:"required"`
	}
	var ids R
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.AdminDepartment.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}
