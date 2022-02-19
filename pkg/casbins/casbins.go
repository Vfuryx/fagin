package casbins

import (
	"fagin/pkg/db"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// casbinModel 模型配置
var casbinModel model.Model

// InitModel 初始化
func InitModel() {
	casbinModel = model.NewModel()

	// m := conf.GetStringMapStringSlice("casbin.model", map[string][]string{})
	// mLen := len(m)
	// for i := 0; i < mLen; i++ {
	//	c := m[strconv.Itoa(i)]
	//	casbinModel.AddDef(c[0], c[1], c[2])
	// }

	//// RESTful 模型 参考 https://casbin.org/docs/zh-CN/supported-models
	// casbin.AddDef("r", "r", "sub, obj, act")
	// casbin.AddDef("p", "p", "sub, obj, act")
	// casbin.AddDef("e", "e", "some(where (p.eft == allow))")
	// casbin.AddDef("m", "m", "r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)")

	// rbac_model.conf 参考 https://github.com/casbin/casbin/blob/master/examples/rbac_model.conf
	casbinModel.AddDef("r", "r", "sub, obj, act")
	casbinModel.AddDef("p", "p", "sub, obj, act")
	casbinModel.AddDef("g", "g", "_, _")
	casbinModel.AddDef("e", "e", "some(where (p.eft == allow))")
	casbinModel.AddDef("m", "m", "r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)")
}

// CasbinService 权限服务
type CasbinService struct {
	E *casbin.Enforcer
}

// CasbinModel 权限结构
type CasbinModel struct {
	gorm.Model
	PType    string `gorm:"type:varchar(64);not null;default:'';comment:'类型';column:p_type;"`
	RoleName string `gorm:"type:varchar(64);not null;default:'';comment:'角色名称';column:role_name;"`
	Path     string `gorm:"type:varchar(64);not null;default:'';comment:'路径';column:path;"`
	Method   string `gorm:"type:varchar(64);not null;default:'';comment:'方法';column:method;"`
}

// Casbin Casbin
var Casbin *CasbinService

// Init 初始化
func Init() {
	InitModel()

	var err error

	Casbin, err = NewCanbin()
	if err != nil {
		panic(err)
	}
}

// NewCanbin 持久化到数据库
func NewCanbin() (*CasbinService, error) {
	// 链接数据库
	a, err := gormadapter.NewAdapterByDB(db.ORM()) // Your driver and data source.
	if err != nil {
		return nil, err
	}
	// 绑定
	e, err := casbin.NewEnforcer(casbinModel, a)
	if err != nil {
		return nil, err
	}
	// 加载
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}

	// 绑定对象
	c := new(CasbinService)
	c.E = e

	return c, nil
}

// 获取 casbin
func (c *CasbinService) casbin() (*casbin.Enforcer, error) {
	// 单例
	if c.E != nil {
		return c.E, nil
	}

	var mx sync.Mutex
	// 防止并发出错
	mx.Lock()
	casbinObj, err := NewCanbin()
	mx.Unlock()

	if err != nil {
		return nil, err
	}

	c.E = casbinObj.E

	return c.E, nil
}

// AddUserRole 添加用户角色
func (c *CasbinService) AddUserRole(userID, role string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.AddGroupingPolicy(userID, role)
}

// DeleteRoleForUser 删除用户的指定角色
func (c *CasbinService) DeleteRoleForUser(userID, role string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.DeleteRoleForUser(userID, role)
}

// GetRolesForUser 获取用户的所有权限
func (c *CasbinService) GetRolesForUser(userID string) ([]string, error) {
	cb, err := c.casbin()
	if err != nil {
		return nil, err
	}

	return cb.GetRolesForUser(userID)
}

// DeleteRolesForUser 删除用户所有角色
func (c *CasbinService) DeleteRolesForUser(userID string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.DeleteRolesForUser(userID)
}

// DeleteRolesForUsers 根据用户删除角色
func (c *CasbinService) DeleteRolesForUsers(userIDs []string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	for i := range userIDs {
		ok, err := cb.DeleteRolesForUser(userIDs[i])
		if err != nil {
			return ok, err
		}
	}

	return true, nil
}

// AddPolicyForUser 添加用户权限
func (c *CasbinService) AddPolicyForUser(userID, obj, act string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.AddPermissionForUser(userID, obj, act)
}

// AddPolicyForRole 添加角色权限
func (c *CasbinService) AddPolicyForRole(role, obj, act string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.AddPolicy(role, obj, act)
}

// AddPolicies 添加策略
func (c *CasbinService) AddPolicies(rules [][]string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.AddPolicies(rules)
}

// AddCasbin 添加角色权限
func (c *CasbinService) AddCasbin(cm *CasbinModel) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.AddPolicy(cm.RoleName, cm.Path, cm.Method)
}

// RemovePolicy 删除权限
func (c *CasbinService) RemovePolicy(role, obj, act string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.RemovePolicy(role, obj, act)
}

// ReCasbin 删除权限
func (c *CasbinService) ReCasbin(cm *CasbinModel) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.RemovePolicy(cm.RoleName, cm.Path, cm.Method)
}

// DeleteRole 删除角色以及权限
func (c *CasbinService) DeleteRole(role string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.DeleteRole(role)
}

// DeletePoliciesByRole 根据角色删除权限
func (c *CasbinService) DeletePoliciesByRole(role string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.RemoveFilteredPolicy(0, role)
}

// CheckRole 验证角色权限
func (c *CasbinService) CheckRole(sub, obj, act string) (bool, error) {
	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	return cb.Enforce(sub, obj, act)
}

// CheckRoles 验证所有角色权限
func (c *CasbinService) CheckRoles(roles []string, obj, act string) (bool, error) {
	l := len(roles)
	if l <= 0 {
		return false, nil
	}

	cb, err := c.casbin()
	if err != nil {
		return false, err
	}

	for i := range roles {
		ok, err := cb.Enforce(roles[i], obj, act)
		if ok && err == nil {
			return true, nil
		}
	}

	return false, nil
}

// GetRoles	获取所有角色
func (c *CasbinService) GetRoles(role string) ([]string, error) {
	cb, err := c.casbin()
	if err != nil {
		return nil, err
	}

	return cb.GetRolesForUser(role)
}
