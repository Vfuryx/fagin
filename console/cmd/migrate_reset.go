package cmd

import (
	"fagin/pkg/db"
	"fagin/pkg/migrate"
	"github.com/spf13/cobra"
)

var migrateResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "重置数据迁移",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Reset()
		// 关闭orm
		defer db.Close()
	},
}

func init() {
	migrateCmd.AddCommand(migrateResetCmd)
}
