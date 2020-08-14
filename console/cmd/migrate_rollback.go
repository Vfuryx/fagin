package cmd

import (
	"fagin/pkg/db"
	"fagin/pkg/migrate"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "回滚到上一步",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Rollback()
		// 关闭orm
		defer db.Close()
	},
}

func init() {
	migrateCmd.AddCommand(rollbackCmd)
}
