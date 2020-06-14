package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/requests/admin"
	"fagin/app/responses/admin_response"
	"fagin/app/service"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"time"
)

type menuController struct{}

var MenuController menuController

func (menuController) Index(ctx *gin.Context) {
	var r admin_request.AdminMenuList
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	params := gin.H{
		"sort": "sort desc, id asc",
	}

	if r.Title != "" {
		params["like_title"] = "%" + r.Title + "%"
	}
	if r.Visible != nil {
		params["visible"] = *r.Visible
	}

	columns := []string{
		"id", "parent_id", "paths", "title", "icon", "type",
		"path", "component", "action", "permission", "sort", "visible",
		"is_link", "created_at",
	}

	menus, err := service.AdminMenuService.Index(params, columns, nil)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrAdminMenuList, nil)
		return
	}

	data := admin_response.AdminMenusList(menus...).Collection()

	app.JsonResponse(ctx, errno.OK, data)
	return
}

func (menuController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	columns := []string{"*"}
	m, err := service.AdminMenuService.Show(id, columns)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrAddAdminMenu, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{
		"id":         m.ID,
		"parent_id":  m.ParentId,
		"paths":      m.Paths,
		"name":       m.Name,
		"title":      m.Title,
		"icon":       m.Icon,
		"type":       m.Type,
		"path":       m.Path,
		"component":  m.Component,
		"action":     m.Action,
		"permission": m.Permission,
		"sort":       m.Sort,
		"visible":    m.Visible,
		"is_link":    m.IsLink,
		"created_at": m.CreatedAt.Format(time.RFC3339),
	})
	return
}

func (menuController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.AdminMenuService.Delete(id)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrDeleteAdminMenu, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (menuController) Store(ctx *gin.Context) {
	var r admin_request.CreateAdminMenu
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	b := admin_menu.AdminMenu{
		ParentId:   r.ParentId,
		Name:       r.Name,
		Title:      r.Title,
		Icon:       r.Icon,
		Type:       *r.Type,
		Path:       r.Path,
		Component:  r.Component,
		Action:     r.Action,
		Permission: r.Permission,
		Sort:       r.Sort,
		Visible:    *r.Visible,
		IsLink:     *r.IsLink,
	}

	spew.Dump(r, b)
	err := service.AdminMenuService.Create(&b)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrAddAdminMenu, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (menuController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r admin_request.UpdateAdminMenu
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	if id == r.ParentId {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	data := map[string]interface{}{
		"parent_id":   r.ParentId,
		"name":       r.Name,
		"title":      r.Title,
		"icon":       r.Icon,
		"type":       *r.Type,
		"path":       r.Path,
		"component":  r.Component,
		"action":     r.Action,
		"permission": r.Permission,
		"sort":       r.Sort,
		"visible":    *r.Visible,
		"is_link":     *r.IsLink,
	}
	err = service.AdminMenuService.Update(id, data)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrUpdateAdminMenu, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}
