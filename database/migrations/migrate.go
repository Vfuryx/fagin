package migrations

import (
	"fagin/app/models"
	"gopkg.in/gormigrate.v1"
	"log"
)

var ORM = models.ORM

/**
 * 加载迁移
 * 格式 m_20xx_xx_xx_xxxxx_action 要严格按照执行顺序
 */
var migrations = []*gormigrate.Migration{
	m_2018_01_23_00001_create_user,
	m_2018_02_01_00001_add_user_col_phone_,
	m_2018_02_01_00002_delete_user_col_phone,

	//...
}

var migrate = gormigrate.New(ORM, gormigrate.DefaultOptions, migrations)

/*
 * 迁移数据库
 */
func Init()  {
	var err error

	if err = migrate.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}

/**
 * 重置最后一条迁移
 */
func Rollback() {
	var err error

	if err = migrate.RollbackLast(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}

/**
 * 重置数据库
 */
func Reset()  {
	var err error

	for {
		if err = migrate.RollbackLast(); err != nil {
			break
		}
	}

	log.Printf("Migration did run successfully")
}