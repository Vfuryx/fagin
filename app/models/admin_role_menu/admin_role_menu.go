package admin_role_menu

type AdminRoleMenu struct {
	ID     uint `gorm:"primarykey"`
	RoleID uint `gorm:"type:int(11) unsigned;not null;default:0;comment:'角色id';column:role_id;"`
	MenuID uint `gorm:"type:int(11) unsigned;not null;default:0;comment:'菜单id';column:menu_id;"`
}
