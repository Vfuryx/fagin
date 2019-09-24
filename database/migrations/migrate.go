package migrations

import (
	"fagin/pkg/db"
	"gopkg.in/gormigrate.v1"
	"log"
	"os"
	"time"
)

/**
 * 迁移集合
 */
var migrations []*gormigrate.Migration

// 确保这里的init是最后一个执行
func migration() *gormigrate.Gormigrate {
	return gormigrate.New(db.ORM, gormigrate.DefaultOptions, migrations)
}

/*
 * 生成迁移文件
 */
func Create(name string)  {
	if name == "" {
		panic("文件名不能为空")
	}

	datetime := time.Now().Format("m_2006_01_02_150405_")

	fileName := datetime + name
	filePath := "./database/migrations/" + fileName + ".go"

	//os.Stat获取文件信息
	if _, err := os.Stat(filePath); err == nil {
		panic("文件已存在")
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	template := "package migrations \n\nimport (\n\t\"github.com/jinzhu/gorm\"\n\t\"gopkg.in/gormigrate.v1\"\n)\n\n" +
		"func init() {\n\tmigrations = append(migrations, &gormigrate.Migration{\n\t\tID: \""+ fileName +"\",\n\t\t" +
		"Migrate: func(tx *gorm.DB) error {\n\t\t\ttype table struct {\n\t\t\t\tgorm.Model\n\t\t\t}\n\t\t\treturn " +
		"tx.AutoMigrate(&table{}).Error\n\t\t},\n\t\tRollback: func(tx *gorm.DB) error {\n\t\t\t" +
		"return tx.DropTableIfExists(\"table\").Error\n\t\t},\n\t})\n}"

	if _, err = file.WriteString(template); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}
}

/*
 * 迁移数据库
 */
func Init() {
	var err error

	if err = migration().Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")

}

/**
 * 重置最后一条迁移
 */
func Rollback() {
	var err error

	if err = migration().RollbackLast(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}

/**
 * 重置数据库
 */
func Reset() {
	var err error

	for {
		if err = migration().RollbackLast(); err != nil {
			break
		}
	}

	log.Printf("Migration did run successfully")
}
