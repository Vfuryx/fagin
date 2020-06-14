package admin_role_menu

type AdminRoleMenu struct {
	RoleID uint `gorm:"type:int(10);not null;default:0;comment:'角色id';column:role_id;"`
	MenuID uint `gorm:"type:int(10);not null;default:0;comment:'菜单id';column:menu_id;"`
}
