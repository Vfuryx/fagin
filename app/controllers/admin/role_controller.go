package admin

import (
	"errors"
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type roleController struct {
	BaseController
}

// RoleController 角色控制器
var RoleController roleController

func (ctr *roleController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, DefaultLimit)

	var r = adminRequest.NewAdminRoleList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
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

	var columns = []string{"*"}

	with := gin.H{"Menus": nil}

	roles, err := services.AdminRoleService.Index(params, columns, with, paginator)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewAdminRoleList(roles...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"items": data,
		"total": paginator.TotalCount,
	})
}

func (ctr *roleController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}

	r, err := services.AdminRoleService.Show(id, columns)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	var menuIDs = new([]uint)
	_ = getMenuTree(r.Menus, 0, menuIDs)

	ctr.ResponseJSONOK(ctx, gin.H{
		"id":         r.ID,
		"type":       r.Type,
		"name":       r.Name,
		"key":        r.Key,
		"sort":       r.Sort,
		"status":     r.Status,
		"created_at": app.TimeToStr(r.CreatedAt),
		"remark":     r.Remark,
		"menu_ids":   menuIDs,
	})
}

func getMenuTree(data []*admin_menu.AdminMenu, pid uint, menuIDs *[]uint) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)

	for index := range data {
		if data[index].ParentID == pid {
			m := map[string]interface{}{
				"id": data[index].ID,
			}

			if children := getMenuTree(data, data[index].ID, menuIDs); len(children) == 0 {
				*menuIDs = append(*menuIDs, data[index].ID)
			}

			result = append(result, m)
		}
	}

	return result
}

func (ctr *roleController) Store(ctx *gin.Context) {
	var r = adminRequest.NewCreateAdminRole()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	ok := services.AdminRoleService.KeyExist(0, r.Key)
	if ok {
		ctr.ResponseJSONErr(ctx, errno.CtxRoleKeyExistErr, nil)
		return
	}

	// 获取菜单组
	var menus []*admin_menu.AdminMenu

	err := admin_menu.NewDao().
		Query(gin.H{"in_id": r.MenuIDs}, []string{"*"}, nil).
		Find(&menus)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	role := admin_role.AdminRole{
		Name: r.Name,
		//Type:   r.Type,
		Key:    r.Key,
		Remark: r.Remark,
		Sort:   r.Sort,
		Status: *r.Status,
		Menus:  menus,
	}

	err = services.AdminRoleService.Create(&role)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *roleController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = adminRequest.NewUpdateAdminRole()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	data := map[string]interface{}{
		"name":    r.Name,
		"type":    r.Type,
		"sort":    r.Sort,
		"status":  *r.Status,
		"remark":  r.Remark,
		"menuIDs": r.MenuIDs,
	}

	err = services.AdminRoleService.Update(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	// 清除角色关联菜单缓存
	err = services.AdminRoleService.RemoveRoleMenusCache(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *roleController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil || id == 1 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.AdminRoleService.Delete(id)
	if err != nil {
		if errors.Is(err, errno.SerRoleRelationExistErr) {
			ctr.ResponseJSONErr(ctx, errno.SerRoleRelationExistErr, nil)
			return
		}

		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)

		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

type updateAdminRoleStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (ctr *roleController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	var r updateAdminRoleStatus

	err = ctx.ShouldBind(&r)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = services.AdminRoleService.UpdateStatus(id, s)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *roleController) Roles(ctx *gin.Context) {
	params := gin.H{
		"orderBy": "sort desc",
	}
	columns := []string{"id", "name"}

	roles, err := services.AdminRoleService.List(params, columns, nil)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewAdminSelectRoleList(roles...).Collection()

	ctr.ResponseJSONOK(ctx, data)
}

func (ctr *roleController) KeyExist(ctx *gin.Context) {
	var r = adminRequest.NewRoleKeyExistRequest()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	ok := services.AdminRoleService.KeyExist(0, r.Key)
	if !ok {
		ctr.ResponseJSONErr(ctx, errno.CtxRoleKeyExistErr, nil)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}
