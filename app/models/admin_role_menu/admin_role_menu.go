package admin_role_menu

type AdminRoleMenu struct {
	ID uint `gorm:"primarykey"`

	RoleID uint `gorm:"uniqueIndex:udx_role_id_menu_id,priority:1;not null;default:0;comment:角色id;column:role_id;"`
	MenuID uint `gorm:"uniqueIndex:udx_role_id_menu_id,priority:2;index;not null;default:0;comment:菜单id;column:menu_id;"`
}
