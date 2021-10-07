package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type menuController struct {
	BaseController
}

// MenuController 菜单控制器
var MenuController menuController

func (mc *menuController) Index(ctx *gin.Context) {
	var r = admin_request.NewAdminMenuList()
	if data, ok := r.Validate(ctx); !ok {
		mc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	params := gin.H{
		"orderBy": "sort desc, id asc",
	}

	if r.Title != "" {
		params["like_title"] = "%" + r.Title + "%"
	}
	if r.Status != nil {
		params["status"] = *r.Status
	}

	columns := []string{
		"id", "parent_id", "icon", "title", "permission", "path",
		"component", "sort", "status", "created_at",
	}

	menus, err := service.AdminMenuService.Index(params, columns, nil)
	if err != nil {
		mc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.AdminMenusList(menus...).Collection()

	mc.ResponseJsonOK(ctx, data)
	return
}

func (mc *menuController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		mc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}
	m, err := service.AdminMenuService.Show(id, columns)
	if err != nil {
		mc.ResponseJsonErrLog(ctx, errno.CtxShowErr, err, nil)
		return
	}

	mc.ResponseJsonOK(ctx, gin.H{
		"id":             m.ID,
		"parent_id":      m.ParentID,
		"paths":          m.Paths,
		"type":           m.Type,
		"name":           m.Name,
		"title":          m.Title,
		"icon":           m.Icon,
		"path":           m.Path,
		"method":         m.Method,
		"component":      m.Component,
		"permission":     m.Permission,
		"is_show":        m.IsShow,
		"redirect":       m.Redirect,
		"frame_src":      m.FrameSrc,
		"current_active": m.CurrentActive,
		"is_hide_child":  m.IsHideChild,
		"sort":           m.Sort,
		"status":         m.Status,
		"created_at":     app.TimeToStr(m.CreatedAt),
	})
	return
}

func (mc *menuController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateAdminMenu()
	if data, ok := r.Validate(ctx); !ok {
		mc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	b := admin_menu.AdminMenu{
		ParentID:      r.ParentId,
		Type:          r.Type,
		Component:     r.Component,
		Name:          r.Name,
		Title:         r.Title,
		Icon:          r.Icon,
		Path:          r.Path,
		Method:        r.Method,
		Permission:    r.Permission,
		Sort:          r.Sort,
		IsShow:        *r.IsShow,
		Redirect:      r.Redirect,
		FrameSrc:      r.FrameSrc,
		CurrentActive: r.CurrentActive,
		Status:        *r.Status,
		IsHideChild:   *r.IsHideChild,
		IsNoCache:     *r.IsNoCache,
	}

	err := service.AdminMenuService.Create(&b)
	if err != nil {
		mc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	mc.ResponseJsonOK(ctx, nil)
	return
}

func (mc *menuController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		mc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewUpdateAdminMenu()
	if data, ok := r.Validate(ctx); !ok {
		mc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	if id == r.ParentId {
		mc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	data := map[string]interface{}{
		"parent_id":      r.ParentId,
		"name":           r.Name,
		"component":      r.Component,
		"title":          r.Title,
		"icon":           r.Icon,
		"type":           r.Type,
		"path":           r.Path,
		"method":         r.Method,
		"permission":     r.Permission,
		"sort":           r.Sort,
		"is_show":        *r.IsShow,
		"status":         *r.Status,
		"frame_src":      r.FrameSrc,
		"current_active": r.CurrentActive,
		"redirect":       r.Redirect,
		"is_hide_child":  r.IsHideChild,
	}
	err = service.AdminMenuService.Update(id, data)
	if err != nil {
		mc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}
	err = service.AdminMenuService.RemoveUserMenusCache(id)
	if err != nil {
		mc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	mc.ResponseJsonOK(ctx, nil)
	return
}

func (mc *menuController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		mc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.AdminMenuService.Delete(id)
	if err != nil {
		if err == errno.SerMenuSubExistErr {
			mc.ResponseJsonErr(ctx, errno.SerMenuSubExistErr, nil)
			return
		}
		if err == errno.SerMenuRelationExistErr {
			mc.ResponseJsonErr(ctx, errno.SerMenuRelationExistErr, nil)
			return
		}
		mc.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err, nil)
		return
	}

	mc.ResponseJsonOK(ctx, nil)
	return
}

func (mc *menuController) All(ctx *gin.Context) {
	var r = admin_request.NewAdminMenuList()
	if data, ok := r.Validate(ctx); !ok {
		mc.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	params := gin.H{
		"orderBy": "sort desc, id asc",
	}

	columns := []string{
		"id", "parent_id", "icon", "title",
	}
	groups, err := service.AdminMenuService.All(params, columns)
	if err != nil {
		mc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.AdminMenusAll(*groups...).Collection()

	mc.ResponseJsonOK(ctx, data)
	return
}
