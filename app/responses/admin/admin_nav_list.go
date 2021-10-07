package admin_responses

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/response"
)

type adminNavList struct {
	ms []admin_menu.AdminMenu
	response.Collect
}

var _ response.IResponse = &adminNavList{}

func AdminNavList(models ...admin_menu.AdminMenu) *adminNavList {
	res := adminNavList{ms: models}
	res.SetCollect(&res)
	return &res
}

func (r *adminNavList) Serialize() []map[string]interface{} {
	return r.RouteTree(r.ms, 0)
}

func (r adminNavList) RouteTree(data []admin_menu.AdminMenu, pid uint) []map[string]interface{} {
	res := make([]map[string]interface{}, 0, 10)
	for index := range data {
		if data[index].ParentID == pid {

			meta := map[string]interface{}{
				"title":              data[index].Title,
				"hideMenu":           data[index].IsShow == 0,
				"hideChildrenInMenu": data[index].IsHideChild != 0,
				"ignoreKeepAlive":    data[index].IsNoCache != 0,
				"orderNo":            data[index].Sort,
				"frameSrc":           data[index].FrameSrc,
				"currentActiveMenu":  data[index].CurrentActive,
				"ignoreRoute":        false, // 默认不忽略路由
			}
			if data[index].Icon != "" {
				meta["icon"] = data[index].Icon
			}
			m := map[string]interface{}{
				"id":        data[index].ID,
				"parentID":  data[index].ParentID,
				"name":      data[index].Name,
				"path":      data[index].Path,
				"component": data[index].Component,
				"redirect":  data[index].Redirect,
				"status":    data[index].Status,
				"meta":      meta,
			}
			if children := r.RouteTree(data, data[index].ID); len(children) > 0 {
				m["children"] = children
			}
			res = append(res, m)
		}
	}
	return res
}
