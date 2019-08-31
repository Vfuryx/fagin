package permissions

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 权限模型(路由)
type Permissions struct {
	gorm.Model
	Name    string `gorm:"type:varchar(64);not null;default:'';comment:'名称';column:name;"`
	Path    string `gorm:"type:varchar(64);not null;default:'';comment:'路径';column:path;"`
	Action  string `gorm:"type:varchar(64);not null;default:'';comment:'方法';column:action;"`
	Comment string `gorm:"type:varchar(64);not null;default:'';comment:'注释';column:comment;"`
	Sort    uint64 `gorm:"type:int(4) unsigned;not null;default:0;comment:'排序 值越大越靠前"`
	Status  uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
}

// 输出格式
type PermissionsSerializer struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Action    string     `json:"action"`
	Comment   string     `json:"comment"`
	Sort      uint64     `json:"sort"`
	Status    uint8      `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

// 序列化接口
func (this *Permissions) Serializer() PermissionsSerializer {
	return PermissionsSerializer{
		ID:        this.ID,
		Name:      this.Name,
		Path:      this.Path,
		Action:    this.Action,
		Comment:   this.Comment,
		Sort:      this.Sort,
		Status:    this.Status,
		CreatedAt: this.CreatedAt,
		UpdatedAt: this.UpdatedAt,
		DeletedAt: this.DeletedAt,
	}
}
