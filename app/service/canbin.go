package service

import (
	"fagin/config"
	"fagin/pkg/db"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
)

// 权限服务
type canbinService struct {
	E *casbin.Enforcer
}

// 权限结构
type CasbinModel struct {
	gorm.Model
	PType    string `gorm:"type:varchar(64);not null;default:'';comment:'类型';column:p_type;"`
	RoleName string `gorm:"type:varchar(64);not null;default:'';comment:'角色名称';column:role_name;"`
	Path     string `gorm:"type:varchar(64);not null;default:'';comment:'路径';column:path;"`
	Method   string `gorm:"type:varchar(64);not null;default:'';comment:'方法';column:method;"`
}

var Canbin canbinService

//持久化到数据库
func (this *canbinService) canbin() *casbin.Enforcer {

	if this.E != nil {
		return this.E
	}
	// 链接数据库
	a := gormadapter.NewAdapterByDB(db.ORM) // Your driver and data source.
	// 绑定
	e := casbin.NewEnforcer(config.Casbin, a)

	// 加载
	e.LoadPolicy()
	return e
}

// 添加用户角色
func (this *canbinService) AddUserRole(userId string, role string) (bool, error) {
	return this.canbin().AddGroupingPolicySafe(userId, role)
}

// 删除用户的指定角色
func (this *canbinService) DeleteRoleForUser(userId string, role string) bool {
	return this.canbin().DeleteRoleForUser(userId, role)
}

// 获取用户的所有权限
func (this *canbinService) GetRolesForUser(userId string) []string {
	roles, err := this.canbin().GetRolesForUser(userId)
	if err != nil {
		return []string{}
	}
	return roles
}

// 删除用户所有角色
func (this *canbinService) DeleteRolesForUser(userId string) bool {
	return this.canbin().DeleteRolesForUser(userId)
}

// 添加用户权限
func (this *canbinService) AddPolicyForUser(userId string, obj string, act string) bool {
	return this.canbin().AddPermissionForUser(userId, obj, act)
}

// 添加角色权限
func (this *canbinService) AddPolicyForRole(role, obj, act string) bool {
	return this.canbin().AddPolicy(role, obj, act)
}

// 添加角色权限
func (this *canbinService) AddCasbin(cm CasbinModel) bool {
	return this.canbin().AddPolicy(cm.RoleName, cm.Path, cm.Method)
}

// 删除权限
func (this *canbinService) ReCasbin(cm CasbinModel) bool {
	return this.canbin().RemovePolicy(cm.RoleName, cm.Path, cm.Method)
}

// 验证角色权限
func (this *canbinService) CheckRole(sub, obj, act string) (bool, error) {
	return this.canbin().EnforceSafe(sub, obj, act)
}

// 验证所有角色权限
func (this *canbinService) CheckRoles(roles []string, obj, act string) bool {
	for _, v := range roles {
		ok, _ := this.canbin().EnforceSafe(v, obj, act)
		if ok {
			return ok
		}
	}
	return false
}

// 	获取所有角色
func (this *canbinService) GetRole() []string {
	roles, err := this.canbin().GetRolesForUser("admin")
	if err != nil {
		return []string{}
	}
	return roles
}

//
//func (this *canbin) T() bool {
//	return this.canbin().EnforceSafe()
//}

// 修改权限
//func (this *canbin) EditCasbin(cm CasbinModel) bool {
//	return this.Canbin().(cm.RoleName, cm.Path, cm.Method)
//}
