package migrate

import (
	"fagin/config"
	_ "fagin/database/migrations" // 载入迁移数据
	"fagin/pkg/db"
	"fagin/pkg/migrate/migration"
	"fmt"
	"gopkg.in/gormigrate.v1"
	"log"
	"os"
	"time"
)

// 确保这里的init是最后一个执行
func migrate() *gormigrate.Gormigrate {
	return gormigrate.New(db.ORM, gormigrate.DefaultOptions, migration.Migrations)
}

// 生成迁移文件
func Create(name string) {
	if name == "" {
		panic("文件名不能为空")
	}

	datetime := time.Now().Format("m_2006_01_02_150405_")

	fileName := datetime + name
	filePath := config.App.MigrationsPath + "/" + fileName + ".go"

	//os.Stat获取文件信息
	if _, err := os.Stat(filePath); err == nil {
		panic("文件已存在")
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	const template = `package migrations

import (
	"%s/pkg/migrate/migration"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	// 自定义表
	type table struct {
		//...
	}
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "%s",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&table{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(&table{}).Error
		},
	})
}
`
	t := fmt.Sprintf(template, config.App.Name, fileName)
	if _, err = file.WriteString(t); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("Migration create run successfully")
}

// 迁移数据库
func Init() {
	var err error

	if err = migrate().Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")

}

// 重置最后一条迁移
func Rollback() {
	var err error

	if err = migrate().RollbackLast(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration rollback run successfully")
}

// 重置数据库
func Reset() {
	var err error

	for {
		if err = migrate().RollbackLast(); err != nil {
			break
		}
	}

	log.Printf("Migration reset run successfully")
}
