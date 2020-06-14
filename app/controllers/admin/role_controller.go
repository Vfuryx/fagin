package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/app/requests/admin"
	"fagin/app/responses/admin_response"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"time"
)

type roleController struct{}

var RoleController roleController

func (roleController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	var r admin_request.AdminRoleList
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	params := gin.H{}
	if r.Name != "" {
		params["like_name"] = "%" + r.Name + "%"
	}
	if r.Key != "" {
		params["like_key"] = "%" + r.Key + "%"
	}
	if r.Status != nil {
		params["status"] = *r.Status
	}
	columns := []string{"*",}
	with := gin.H{"Menus": nil}

	roles, err := service.AdminRoleService.Index(params, columns, with, &paginator)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrAdminRoleList, nil)
		return
	}

	data := admin_response.AdminRoleList(roles...).Collection()

	app.JsonResponse(ctx, errno.OK, gin.H{
		"roles":     data,
		"paginator": paginator,
	})
	return
}

func (roleController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	columns := []string{"*"}
	r, err := service.AdminRoleService.Show(id, columns)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrAdminRole, nil)
		return
	}

	menuIDs := make([]uint, 0, 20)
	getTreeNode(r.Menus, 0, &menuIDs)

	app.JsonResponse(ctx, errno.OK, gin.H{
		"id":         r.ID,
		"name":       r.Name,
		"key":        r.Key,
		"sort":       r.Sort,
		"status":     r.Status,
		"created_at": r.CreatedAt.Format(time.RFC3339),
		"remark":     r.Remark,
		"menu_ids":   menuIDs,
	})
	return
}

// 获取没有子节点的的菜单
func getTreeNode(menus []admin_menu.AdminMenu, pID uint, node *[]uint) []uint {
	var m = make([]uint, 0, 10)
	for _, menu := range menus {
		if menu.ParentId == pID {
			m = append(m, menu.ID)
			mm := getTreeNode(menus, menu.ID, node)
			if len(mm) <= 0 {
				*node = append(*node, menu.ID)
			}
		}
	}
	return m
}

func (roleController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.AdminRoleService.Delete(id)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrDeleteAdminRole, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

type roleIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (roleController) Dels(ctx *gin.Context) {
	var ids roleIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.AdminRoleService.Deletes(ids.IDs)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrDeleteAdminRole, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (roleController) Store(ctx *gin.Context) {
	var r admin_request.CreateAdminRole
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}
	spew.Dump(r.MenuIDs)

	var menus []admin_menu.AdminMenu
	err := admin_menu.Dao().Query(gin.H{"in_id": r.MenuIDs}, []string{"*"}, nil).Find(&menus)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrAddAdminRole, nil)
		return
	}

	role := admin_role.AdminRole{
		Name:   r.Name,
		Key:    r.Key,
		Remark: r.Remark,
		Sort:   r.Sort,
		Status: *r.Status,
		Menus:  menus,
	}

	err = service.AdminRoleService.Create(&role)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrAddAdminMenu, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (roleController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r admin_request.UpdateAdminRole
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	data := map[string]interface{}{
		"sort":    r.Sort,
		"status":  *r.Status,
		"remark":  r.Remark,
		"menuIDs": r.MenuIDs,
	}
	err = service.AdminRoleService.Update(id, data)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrUpdateAdminRole, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

type updateAdminRoleStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (roleController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r updateAdminRoleStatus
	err = ctx.ShouldBind(&r)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}
	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = service.AdminRoleService.UpdateStatus(id, s)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrUpdateAdminRole, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (roleController) Roles(ctx *gin.Context) {

	params := gin.H{
		"sort": "sort desc",
	}
	columns := []string{"id", "name",}

	roles, err := service.AdminRoleService.List(params, columns, nil)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrAdminRoleList, nil)
		return
	}

	data := admin_response.AdminSelectRoleList(roles...).Collection()

	app.JsonResponse(ctx, errno.OK, gin.H{
		"roles": data,
	})
	return
}
