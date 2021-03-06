package cmd

import (
	"fagin/pkg/migrate"
	"github.com/spf13/cobra"
)

var MigrateRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "回滚到上一步",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Rollback()
	},
}

