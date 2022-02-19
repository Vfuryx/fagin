package responses

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminNavList struct {
	ms []*admin_menu.AdminMenu

	response.Collect
}

func NewAdminNavList(models ...*admin_menu.AdminMenu) response.Response {
	res := adminNavList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (r *adminNavList) Serialize() []map[string]interface{} {
	var menu *admin_menu.AdminMenu

	res := make([]map[string]interface{}, 0, response.DefCap)

	for index := range r.ms {
		menu = r.ms[index]

		meta := map[string]interface{}{
			"title":              menu.Title,
			"hideMenu":           menu.IsShow == 0,
			"hideChildrenInMenu": menu.IsHideChild != 0,
			"ignoreKeepAlive":    menu.IsNoCache != 0,
			"orderNo":            menu.Sort,
			"frameSrc":           menu.FrameSrc,
			"currentActiveMenu":  menu.CurrentActive,
			"ignoreRoute":        false, // 默认不忽略路由
		}
		if menu.Icon != "" {
			meta["icon"] = menu.Icon
		}

		m := map[string]interface{}{
			"id":        menu.ID,
			"pid":       menu.ParentID,
			"name":      menu.Name,
			"path":      menu.Path,
			"component": menu.Component,
			"redirect":  menu.Redirect,
			"status":    menu.Status,
			"meta":      meta,
		}

		res = append(res, m)
	}

	return res
}
