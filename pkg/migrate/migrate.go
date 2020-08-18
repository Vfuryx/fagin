package migrate

import (
	"fagin/config"
	"fagin/pkg/db"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

// 确保这里的init是最后一个执行
//func migrate() *gormigrate.Gormigrate {
//	return gormigrate.New(db.ORM, gormigrate.DefaultOptions, migration.Migrations)
//}

//// 生成迁移文件
//func Create(name string) {
//	if name == "" {
//		panic("文件名不能为空")
//	}
//
//	datetime := time.Now().Format("m_2006_01_02_150405_")
//
//	fileName := datetime + name
//	filePath := config.App.MigrationsPath + "/" + fileName + ".go"
//
//	//os.Stat获取文件信息
//	if _, err := os.Stat(filePath); err == nil {
//		panic("文件已存在")
//	}
//
//	file, err := os.Create(filePath)
//	if err != nil {
//		panic(err)
//	}
//
//	const template = `package migrations
//
//import (
//	"%s/pkg/migrate/migration"
//	"gorm.io/gorm"
//	"gopkg.in/gormigrate.v1"
//)
//
//func init() {
//	// 自定义表
//	type table struct {
//		//...
//	}
//	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
//		ID: "%s",
//		Migrate: func(tx *gorm.DB) error {
//			return tx.AutoMigrate(&table{}).Error
//		},
//		Rollback: func(tx *gorm.DB) error {
//			return tx.DropTableIfExists(&table{}).Error
//		},
//	})
//}
//`
//	t := fmt.Sprintf(template, config.App.Name, fileName)
//	if _, err = file.WriteString(t); err != nil {
//		panic(err)
//	}
//
//	if err = file.Close(); err != nil {
//		panic(err)
//	}
//
//	log.Printf("Migration create run successfully")
//}

//// 迁移数据库
//func Init() {
//	var err error
//
//	if err = migrate().Migrate(); err != nil {
//		log.Fatalf("Could not migrate: %v", err)
//	}
//
//	log.Printf("Migration did run successfully")
//
//}
//
//// 重置最后一条迁移
//func Rollback() {
//	var err error
//
//	if err = migrate().RollbackLast(); err != nil {
//		log.Fatalf("Could not migrate: %v", err)
//	}
//
//	log.Printf("Migration rollback run successfully")
//}
//
//// 重置数据库
//func Reset() {
//	var err error
//
//	for {
//		if err = migrate().RollbackLast(); err != nil {
//			break
//		}
//	}
//
//	log.Printf("Migration reset run successfully")
//}

func CreateMigration(name string) error {
	dir := config.App.MigrationsPath
	var templateContent = `
-- +migrate Up

-- +migrate Down
`

	tpl := template.Must(template.New("new_migration").Parse(templateContent))

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	}

	fileName := fmt.Sprintf("%s_%s.sql", time.Now().Format("20060102150405"), strings.TrimSpace(name))
	pathName := path.Join(dir, fileName)
	f, err := os.Create(pathName)

	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	if err := tpl.Execute(f, nil); err != nil {
		return err
	}

	fmt.Printf("Created migration %s", pathName)
	return nil
}

// 迁移数据库
func Migration() {
	migrations := &migrate.FileMigrationSource{
		Dir: config.App.MigrationsPath,
	}

	database, err := db.ORM.DB()
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(database, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

// 重置最后一条迁移
func Rollback() {
	source := &migrate.FileMigrationSource{
		Dir: config.App.MigrationsPath,
	}

	database, err := db.ORM.DB()
	if err != nil {
		panic(err)
	}

	_, err = migrate.ExecMax(database, "mysql", source, migrate.Down, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied 1 migration")
}

// 重置数据库
func Reset() {
	migrations := &migrate.FileMigrationSource{
		Dir: config.App.MigrationsPath,
	}

	database, err := db.ORM.DB()
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(database, "mysql", migrations, migrate.Down)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
