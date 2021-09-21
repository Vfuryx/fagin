package admin

import (
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"time"
)

type menuController struct {
	BaseController
}

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
	if r.IsShow != nil {
		params["is_show"] = *r.IsShow
	}

	columns := []string{
		"id", "title", "icon", "type", "path", "sort", "is_show",
		"created_at", "component", "redirect", "target", "status",
		"is_hide_child", "parent_id",
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
		mc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	mc.ResponseJsonOK(ctx, gin.H{
		"id":            m.ID,
		"parent_id":     m.ParentID,
		"paths":         m.Paths,
		"name":          m.Name,
		"title":         m.Title,
		"icon":          m.Icon,
		"type":          m.Type,
		"path":          m.Path,
		"method":        m.Method,
		"component":     m.Component,
		"permission":    m.Permission,
		"sort":          m.Sort,
		"is_show":       m.IsShow,
		"redirect":      m.Redirect,
		"target":        m.Target,
		"status":        m.Status,
		"is_hide_child": m.IsHideChild,
		"created_at":    m.CreatedAt.Format(time.RFC3339),
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
		ParentID:    r.ParentId,
		Component:   r.Component,
		Name:        r.Name,
		Title:       r.Title,
		Icon:        r.Icon,
		Type:        r.Type,
		Path:        r.Path,
		Method:      r.Method,
		Permission:  r.Permission,
		Sort:        r.Sort,
		IsShow:      *r.IsShow,
		Redirect:    r.Redirect,
		Target:      r.Target,
		Status:      *r.Status,
		IsHideChild: *r.IsHideChild,
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
		"parent_id":     r.ParentId,
		"name":          r.Name,
		"component":     r.Component,
		"title":         r.Title,
		"icon":          r.Icon,
		"type":          *r.Type,
		"path":          r.Path,
		"method":        r.Method,
		"permission":    r.Permission,
		"sort":          r.Sort,
		"is_show":       *r.IsShow,
		"status":        *r.Status,
		"target":        r.Target,
		"redirect":      r.Redirect,
		"is_hide_child": r.IsHideChild,
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
		"type":    r.Type,
	}

	columns := []string{"*"}
	groups, err := service.AdminMenuService.All(params, columns)
	if err != nil {
		mc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.AdminMenusList(*groups...).Collection()

	mc.ResponseJsonOK(ctx, data)
	return
}
