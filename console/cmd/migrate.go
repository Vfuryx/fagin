package cmd

import (
	"fagin/pkg/migrate"
	//"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "数据迁移",
	Long:    ``,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Migration()
	},
}
