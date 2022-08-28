package admin

import (
	"errors"
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type menuController struct {
	BaseController
}

// MenuController 菜单控制器
var MenuController menuController

func (ctr *menuController) Index(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.AdminMenuList](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
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

	menus, err := services.AdminMenuService.Index(params, columns, nil)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewAdminMenusList(menus...).Collection()

	ctr.ResponseJSONOK(ctx, data)
}

func (ctr *menuController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}

	m, err := services.AdminMenuService.Show(id, columns)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
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
		"is_no_cache":    m.IsNoCache,
		"sort":           m.Sort,
		"status":         m.Status,
		"created_at":     app.TimeToStr(m.CreatedAt),
	})
}

func (ctr *menuController) Store(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.CreateAdminMenu](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	b := admin_menu.AdminMenu{
		ParentID:      r.ParentID,
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

	err := services.AdminMenuService.Create(&b)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *menuController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	r, msg := request.Validation[adminRequest.UpdateAdminMenu](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	if id == r.ParentID {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	data := map[string]interface{}{
		"parent_id":      r.ParentID,
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
		"is_no_cache":    *r.IsNoCache,
	}

	err = services.AdminMenuService.Update(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	err = services.AdminMenuService.RemoveUserMenusCache(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *menuController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.AdminMenuService.Delete(id)
	if err != nil {
		if errors.Is(err, errno.SerMenuSubExistErr) {
			ctr.ResponseJSONErr(ctx, errno.SerMenuSubExistErr, nil)
			return
		}

		if errors.Is(err, errno.SerMenuRelationExistErr) {
			ctr.ResponseJSONErr(ctx, errno.SerMenuRelationExistErr, nil)
			return
		}

		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)

		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *menuController) All(ctx *gin.Context) {
	_, msg := request.Validation[adminRequest.AdminMenuList](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	params := gin.H{
		"orderBy": "sort desc, id asc",
	}

	columns := []string{
		"id", "parent_id", "icon", "title",
	}

	groups, err := services.AdminMenuService.All(params, columns)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewAdminMenusAll(groups...).Collection()

	ctr.ResponseJSONOK(ctx, data)
}
