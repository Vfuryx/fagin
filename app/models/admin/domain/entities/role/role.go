package role

import "gorm.io/plugin/soft_delete"

type AdminRole struct {
	ID        uint                  `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt uint                  `json:"created_at,omitempty"`
	UpdatedAt uint                  `json:"updated_at,omitempty"`
	DeletedAt soft_delete.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Type   uint8  `gorm:"not null;default:0;comment:菜单类型 0:总后台 1:商家后台 2:集团后台;column:type;"`
	Name   string `gorm:"type:varchar(128);not null;default:'';comment:角色名称;column:name;"`
	Key    string `gorm:"type:varchar(32);not null;default:'';comment:角色关键字;column:key;"`
	Remark string `gorm:"type:varchar(255);not null;default:'';comment:角色备注;column:remark;"`
	Sort   uint   `gorm:"not null;default:0;comment:排序，数字越大越靠前;column:sort"`
	Status uint8  `gorm:"not null;default:1;comment:状态: 0=>关闭 1=>开启;column:status;"`

	//Menus []m.AdminMenu `gorm:"many2many:admin_role_menus;joinForeignKey:role_id;JoinReferences:menu_id;"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (*AdminRole) TableName() string {
	return "admin_role"
}

func (r *AdminRole) Add(username string, nickName string) error {
	var err error
	//if err = user.setUsername(username); err != nil {
	//	return err
	//}
	//if err = user.setNickName(nickName); err != nil {
	//	return err
	//}

	return err
}
