package admin

import (
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/admin_permission"
	"fagin/app/models/admin_permission_group"
	admin_request "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type permissionController struct {
	BaseController
}

var PermissionController permissionController

// GroupPermissions 分组权限列表
func (pc *permissionController) GroupPermissions(ctx *gin.Context) {
	columns := []string{
		"id", "name", "sort", "created_at",
	}
	var r = admin_request.NewPermissionList()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	params := gin.H{
		"orderBy": "sort desc, id asc",
		"type":    r.Type,
	}
	group, err := service.AdminPermissionService.
		GroupPermissions(params, columns, nil)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxListErr, nil, nil)
		return
	}

	data := response.GroupPermissionList(group...).Collection()
	pc.ResponseJsonOK(ctx, data)
	return
}

func (pc *permissionController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	var r = admin_request.NewPermissionList()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	params := gin.H{
		"orderBy": "sort desc, id asc",
		"type":    r.Type,
	}

	if r.Name != "" {
		params["like_name"] = "%" + r.Name + "%"
	}

	columns := []string{
		"id", "gid", "name", "path", "method", "sort",
		"status", "content", "created_at",
	}

	with := gin.H{
		"Group": nil,
	}

	permissions, err := service.AdminPermissionService.
		PermissionList(params, columns, with, paginator)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.PermissionList(permissions...).Collection()

	pc.ResponseJsonOK(ctx, gin.H{
		"list":      data,
		"paginator": paginator,
	})
	return
}

func (pc *permissionController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		pc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}
	m, err := service.AdminPermissionService.PermissionShow(id, columns)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	pc.ResponseJsonOK(ctx, gin.H{
		"id":     m.ID,
		"name":   m.Name,
		"gid":    m.GID,
		"paths":  m.Path,
		"method": m.Method,
		"sort":   m.Sort,
		"status": m.Status,
		"type":   m.Type,
		//"created_at": m.CreatedAt.Format(time_format.Def),
	})
	return
}

func (pc *permissionController) Store(ctx *gin.Context) {
	var r = admin_request.NewUpdatePermission()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	b := admin_permission.AdminPermission{
		Name:   r.Name,
		GID:    r.GID,
		Path:   r.Path,
		Method: r.Method,
		Status: *r.Status,
		Sort:   r.Sort,
		Type:   r.Type,
	}

	err := service.AdminPermissionService.PermissionCreate(&b)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	pc.ResponseJsonOK(ctx, nil)
	return
}

func (pc *permissionController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		pc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewUpdatePermission()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	data := map[string]interface{}{
		"type":   r.Type,
		"name":   r.Name,
		"sort":   r.Sort,
		"gid":    r.GID,
		"path":   r.Path,
		"method": r.Method,
		"status": *r.Status,
	}
	err = service.AdminPermissionService.PermissionUpdate(id, data)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	pc.ResponseJsonOK(ctx, nil)
	return
}

func (pc *permissionController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		pc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.AdminPermissionService.PermissionDelete(id)
	if err != nil {
		if err == errno.SerPermissionRelationExistErr {
			pc.ResponseJsonErr(ctx, errno.SerPermissionRelationExistErr, nil)
			return
		}
		pc.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err, nil)
		return
	}

	pc.ResponseJsonOK(ctx, nil)
	return
}

func (pc *permissionController) GIndex(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	var r = admin_request.NewPermissionGroupList()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	params := gin.H{
		"orderBy": "sort desc, id asc",
	}

	if r.Name != "" {
		params["like_name"] = "%" + r.Name + "%"
	}
	params["type"] = r.Type

	columns := []string{"id", "name", "type", "sort", "created_at"}

	groups, err := service.AdminPermissionService.
		PermissionGroupList(params, columns, nil, paginator)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.PermissionGroupList(groups...).Collection()

	pc.ResponseJsonOK(ctx, gin.H{
		"list":      data,
		"paginator": paginator,
	})
	return
}

func (pc *permissionController) GAll(ctx *gin.Context) {
	var r = admin_request.NewPermissionGroupList()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	params := gin.H{
		"type": r.Type,
	}

	columns := []string{"id", "name", "sort", "created_at"}
	groups, err := service.AdminPermissionService.PermissionGroupAll(params, columns)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.PermissionGroupList(*groups...).Collection()

	pc.ResponseJsonOK(ctx, data)
	return
}

func (pc *permissionController) GShow(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		pc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}
	m, err := service.AdminPermissionService.PermissionGroupShow(id, columns)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxShowErr, err, nil)
		return
	}

	pc.ResponseJsonOK(ctx, gin.H{
		"id":         m.ID,
		"name":       m.Name,
		"sort":       m.Sort,
		"type":       m.Type,
		"created_at": m.CreatedAt.Format(enums.TimeFormatDef.Get()),
	})
	return
}

func (pc *permissionController) GStore(ctx *gin.Context) {
	var r = admin_request.NewUpdatePermissionGroup()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	b := admin_permission_group.AdminPermissionGroup{
		Name: r.Name,
		Sort: r.Sort,
		Type: r.Type,
	}

	err := service.AdminPermissionService.
		PermissionGroupCreate(&b)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	pc.ResponseJsonOK(ctx, nil)
	return
}

func (pc *permissionController) GUpdate(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		pc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewUpdatePermissionGroup()
	if data, ok := r.Validate(ctx); !ok {
		pc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	data := map[string]interface{}{
		"name": r.Name,
		"sort": r.Sort,
		"type": r.Type,
	}
	err = service.AdminPermissionService.
		PermissionGroupUpdate(id, data)
	if err != nil {
		pc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	pc.ResponseJsonOK(ctx, nil)
	return
}
