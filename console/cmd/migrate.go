package cmd

import (
	"fagin/pkg/migrate"
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
