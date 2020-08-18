package cmd

import (
	"fagin/pkg/migrate"
	"github.com/spf13/cobra"
)

var migrateResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "重置数据迁移",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Reset()
	},
}

func init() {
	migrateCmd.AddCommand(migrateResetCmd)
}
