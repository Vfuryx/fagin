package service

import (
	"fagin/config"
	"fagin/pkg/db"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
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
func (c *canbinService) canbin() (*casbin.Enforcer, error) {

	if c.E != nil {
		return c.E, nil
	}
	// 链接数据库
	a, err := gormadapter.NewAdapterByDB(db.ORM) // Your driver and data source.
	if err != nil {
		return nil, err
	}
	// 绑定
	e, err := casbin.NewEnforcer(config.Casbin, a)
	if err != nil {
		return nil, err
	}

	// 加载
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// 添加用户角色
func (c *canbinService) AddUserRole(userId string, role string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.AddGroupingPolicy(userId, role)
}

// 删除用户的指定角色
func (c *canbinService) DeleteRoleForUser(userId string, role string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.DeleteRoleForUser(userId, role)
}

// 获取用户的所有权限
func (c *canbinService) GetRolesForUser(userId string) ([]string, error) {
	cb, err := c.canbin()
	if err != nil {
		return nil, err
	}
	return cb.GetRolesForUser(userId)
}

// 删除用户所有角色
func (c *canbinService) DeleteRolesForUser(userId string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.DeleteRolesForUser(userId)
}

func (c *canbinService) DeleteRolesForUsers(userIDs []string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	for _, u := range userIDs {
		ok, err := cb.DeleteRolesForUser(u)
		if err != nil {
			return ok, err
		}
	}

	return true, nil
}

// 添加用户权限
func (c *canbinService) AddPolicyForUser(userId string, obj string, act string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.AddPermissionForUser(userId, obj, act)
}

// 添加角色权限
func (c *canbinService) AddPolicyForRole(role, obj, act string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.AddPolicy(role, obj, act)
}

func (c *canbinService) AddPolicies(rules [][]string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.AddPolicies(rules)
}

// 添加角色权限
func (c *canbinService) AddCasbin(cm CasbinModel) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.AddPolicy(cm.RoleName, cm.Path, cm.Method)
}

// 删除权限
func (c *canbinService) RemovePolicy(role, obj, act string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.RemovePolicy(role, obj, act)
}

// 删除权限
func (c *canbinService) ReCasbin(cm CasbinModel) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.RemovePolicy(cm.RoleName, cm.Path, cm.Method)
}

// 删除角色以及权限
func (c *canbinService) DeleteRole(role string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.DeleteRole(role)
}

// 根据角色删除权限
func (c *canbinService) DeletePoliciesByRole(role string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.RemoveFilteredPolicy(0, role)
}

// 验证角色权限
func (c *canbinService) CheckRole(sub, obj, act string) (bool, error) {
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	return cb.Enforce(sub, obj, act)
}

// 验证所有角色权限
func (c *canbinService) CheckRoles(roles []string, obj, act string) (bool, error) {
	l := len(roles)
	if l <= 0 {
		return false, nil
	}
	cb, err := c.canbin()
	if err != nil {
		return false, err
	}
	for _, v := range roles {
		ok, err := cb.Enforce(v, obj, act)
		if ok && err == nil {
			return true, nil
		}
	}
	return false, nil
}

// 	获取所有角色
func (c *canbinService) GetRole(role string) ([]string, error) {
	cb, err := c.canbin()
	if err != nil {
		return nil, err
	}

	return cb.GetRolesForUser(role)
}

//
//func (this *canbin) T() bool {
//	return this.canbin().EnforceSafe()
//}

// 修改权限
//func (this *canbin) EditCasbin(cm CasbinModel) bool {
//	return this.Canbin().(cm.RoleName, cm.Path, cm.Method)
//}
