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

func CreateMigration(name string) error {
	dir := config.App().MigrationsPath

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

// Migration 迁移数据库
func Migration() {
	migrations := &migrate.FileMigrationSource{
		Dir: config.App().MigrationsPath,
	}

	database, err := db.ORM().DB()
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(database, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

// Rollback 重置最后一条迁移
func Rollback() {
	source := &migrate.FileMigrationSource{
		Dir: config.App().MigrationsPath,
	}

	database, err := db.ORM().DB()
	if err != nil {
		panic(err)
	}

	_, err = migrate.ExecMax(database, "mysql", source, migrate.Down, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied 1 migration")
}

// Reset 重置数据库
func Reset() {
	migrations := &migrate.FileMigrationSource{
		Dir: config.App().MigrationsPath,
	}

	database, err := db.ORM().DB()
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(database, "mysql", migrations, migrate.Down)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
