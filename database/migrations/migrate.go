package migrations

import (
	"fagin/pkg/db"
	"gopkg.in/gormigrate.v1"
	"log"
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
