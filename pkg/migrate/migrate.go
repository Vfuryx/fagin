package migrate

import (
	"fagin/config"
	"fagin/pkg/db"
	"fagin/utils"
	"path/filepath"

	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
)

var migrations = make([]*gormigrate.Migration, 0, 20)

func newMigrate() *gormigrate.Gormigrate {
	option := &gormigrate.Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}

	return gormigrate.New(db.ORM(), option, migrations)
}

func Add(migration *gormigrate.Migration) {
	migrations = append(migrations, migration)
}

// CreateMigration 创建迁移文件
func CreateMigration(name string) error {
	dir := config.App().MigrationsPath()

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	}

	name = utils.Underscore(strings.TrimSpace(name))
	fileName := fmt.Sprintf("%s_%s.go", time.Now().Format("20060102150405"), name)
	mPath := path.Join(dir, fileName)

	// 判断是否有重复文件
	dirs, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	ext := ""
	for _, f := range dirs {
		if ext = filepath.Ext(f.Name()); ext == ".go" {
			// 去除文件后缀 .go
			n := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			if n[strings.IndexByte(n, '_')+1:] == name {
				log.Fatalln("文件已经存在")
			}
		}
	}

	if err != nil {
		return err
	}

	f, err := os.Create(mPath)

	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()

	var templateContent = `
package migrations

import (
	"fagin/pkg/migrate"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	migrate.Add(&gormigrate.Migration{
		ID: "%s",
		Migrate: func(tx *gorm.DB) error {
			return
		},
		Rollback: func(tx *gorm.DB) error {
			return
		},
	})
}
`

	if _, err = f.WriteString(fmt.Sprintf(templateContent, fileName)); err != nil {
		return err
	}

	fmt.Printf("Created migration %s", mPath)

	return nil
}

// Migration 迁移数据库
func Migration() {
	err := newMigrate().Migrate()
	if err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

// Rollback 重置最后一条迁移
func Rollback() {
	var err error

	if err = newMigrate().RollbackLast(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

// Reset 重置数据库
func Reset() {
	var err error

	for {
		if err = newMigrate().RollbackLast(); err != nil {
			break
		}
	}
}

// Refresh 回滚所有迁移，并运行所有迁移
func Refresh() {
	// 回滚所有迁移
	Reset()

	// 再次执行所有迁移
	Migration()
}
