package admin

import (
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_permission"
	"fagin/app/models/admin_role"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"time"

	"github.com/gin-gonic/gin"
)

type roleController struct {
	BaseController
}

var RoleController roleController

func (rc *roleController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	var r = adminRequest.NewAdminRoleList()
	if data, ok := r.Validate(ctx); !ok {
		rc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	params := gin.H{
		"type": r.Type,
	}
	if r.Name != "" {
		params["like_name"] = "%" + r.Name + "%"
	}
	if r.Key != "" {
		params["like_key"] = "%" + r.Key + "%"
	}
	if r.Status != nil {
		params["status"] = *r.Status
	}
	var columns []string
	with := gin.H{"Menus": nil}

	roles, err := service.AdminRoleService.Index(params, columns, with, paginator)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.AdminRoleList(roles...).Collection()

	rc.ResponseJsonOK(ctx, gin.H{
		"roles":     data,
		"paginator": paginator,
	})
	return
}

func (rc *roleController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		rc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}
	r, err := service.AdminRoleService.Show(id, columns)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxShowErr, err, nil)
		return
	}

	menuIDs := make([]uint, 0, 20)
	for _, menu := range r.Menus {
		menuIDs = append(menuIDs, menu.ID)
	}

	rc.ResponseJsonOK(ctx, gin.H{
		"id":          r.ID,
		"type":        r.Type,
		"name":        r.Name,
		"key":         r.Key,
		"sort":        r.Sort,
		"status":      r.Status,
		"created_at":  r.CreatedAt.Format(time.RFC3339),
		"remark":      r.Remark,
		"menu_ids":    menuIDs,
		"permissions": getPermissions(r.Permissions),
	})
	return
}

func getPermissions(p []admin_permission.AdminPermission) []gin.H {
	ps := make([]gin.H, 0, 20)
	item := gin.H{}
	for _, permission := range p {
		item = gin.H{
			"id":   permission.ID,
			"name": permission.Name,
		}
		ps = append(ps, item)
	}
	return ps
}

func (rc *roleController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil || id == 1 {
		rc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.AdminRoleService.Delete(id)
	if err != nil {
		if err == errno.SerRoleRelationExistErr {
			rc.ResponseJsonErr(ctx, errno.SerRoleRelationExistErr, nil)
			return
		}
		rc.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err, nil)
		return
	}

	rc.ResponseJsonOK(ctx, nil)
	return
}

func (rc *roleController) Store(ctx *gin.Context) {
	var r = adminRequest.NewCreateAdminRole()
	if data, ok := r.Validate(ctx); !ok {
		rc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	ok := service.AdminRoleService.KeyExist(0, r.Key)
	if !ok {
		rc.ResponseJsonErr(ctx, errno.CtxRoleKeyExistErr, nil)
		return
	}

	// 获取菜单组
	var menus []admin_menu.AdminMenu
	err := admin_menu.NewDao().
		Query(gin.H{"in_id": r.MenuIDs}, []string{"*"}, nil).
		Find(&menus)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}
	// 获取权限组
	var permissions []admin_permission.AdminPermission
	err = admin_permission.NewDao().
		Query(gin.H{"in_id": r.PermissionIDs}, []string{"*"}, nil).
		Find(&permissions)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	role := admin_role.AdminRole{
		Name:        r.Name,
		Type:        r.Type,
		Key:         r.Key,
		Remark:      r.Remark,
		Sort:        r.Sort,
		Status:      *r.Status,
		Menus:       menus,
		Permissions: permissions,
	}

	err = service.AdminRoleService.Create(&role)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	rc.ResponseJsonOK(ctx, nil)
	return
}

func (rc *roleController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		rc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = adminRequest.NewUpdateAdminRole()
	if data, ok := r.Validate(ctx); !ok {
		rc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	data := map[string]interface{}{
		"name":           r.Name,
		"type":           r.Type,
		"sort":           r.Sort,
		"status":         *r.Status,
		"remark":         r.Remark,
		"menuIDs":        r.MenuIDs,
		"permissionsIDs": r.PermissionIDs,
	}

	err = service.AdminRoleService.Update(id, data)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	// 清除角色关联菜单缓存
	err = service.AdminRoleService.RemoveRoleMenusCache(id)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	rc.ResponseJsonOK(ctx, nil)
	return
}

type updateAdminRoleStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (rc *roleController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		rc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r updateAdminRoleStatus
	err = ctx.ShouldBind(&r)
	if err != nil {
		rc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}
	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = service.AdminRoleService.UpdateStatus(id, s)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	rc.ResponseJsonOK(ctx, nil)
	return
}

func (rc *roleController) Roles(ctx *gin.Context) {
	params := gin.H{
		"orderBy": "sort desc",
	}
	columns := []string{"id", "name"}

	roles, err := service.AdminRoleService.List(params, columns, nil)
	if err != nil {
		rc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.AdminSelectRoleList(roles...).Collection()

	rc.ResponseJsonOK(ctx, gin.H{
		"roles": data,
	})
	return
}

func (rc *roleController) KeyExist(ctx *gin.Context) {
	var r = adminRequest.NewRoleKeyExistRequest()
	if data, ok := r.Validate(ctx); !ok {
		rc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	ok := service.AdminRoleService.KeyExist(0, r.Key)
	if !ok {
		rc.ResponseJsonErr(ctx, errno.CtxRoleKeyExistErr, nil)
		return
	}

	rc.ResponseJsonOK(ctx, nil)
}
