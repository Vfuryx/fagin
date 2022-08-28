package admin_user_role

// AdminUserRole 管理员角色关联模型
type AdminUserRole struct {
	ID uint `gorm:"primarykey"`

	AdminID uint `gorm:"uniqueIndex:udx_type_key,priority:1;not null;default:0;comment:管理员id;column:admin_id;"`
	RoleID  uint `gorm:"uniqueIndex:udx_type_key,priority:2;index;not null;default:0;comment:角色id;column:role_id;"`
}
