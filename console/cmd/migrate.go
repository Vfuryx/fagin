package cmd

import (
	"fagin/pkg/db"
	"fagin/pkg/migrate"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据迁移",
	Long: ``,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Init()
		// 关闭orm
		defer db.Close()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
