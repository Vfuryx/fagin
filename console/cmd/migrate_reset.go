package cmd

import (
	"fagin/pkg/migrate"
	"github.com/spf13/cobra"
)

// MigrateResetCmd 重置数据迁移
var MigrateResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "重置数据迁移",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Reset()
	},
}
