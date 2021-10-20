package admin_user_role

// AdminUserRole 管理员角色关联模型
type AdminUserRole struct {
	ID      uint `gorm:"primarykey"`
	AdminID uint `gorm:"type:int(11) unsigned;not null;default:0;comment:'管理员id';column:admin_id;"`
	RoleID  uint `gorm:"type:int(11) unsigned;not null;default:0;comment:'角色id';column:role_id;"`
}
