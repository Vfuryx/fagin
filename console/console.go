package console

import (
	"fagin/console/cmd"
	"github.com/spf13/cobra"
)

func Handle(cobra *cobra.Command) {
	cobra.AddCommand(cmd.AdminCmd)
	cobra.AddCommand(cmd.CacheCmd)
	cobra.AddCommand(cmd.ControllerCmd)
	cobra.AddCommand(cmd.KeyCmd)
	cmd.MigrateCmd.AddCommand(cmd.MigrateCreateCmd)
	cmd.MigrateCmd.AddCommand(cmd.MigrateResetCmd)
	cmd.MigrateCmd.AddCommand(cmd.MigrateRollbackCmd)
	cobra.AddCommand(cmd.MigrateCmd)
	cobra.AddCommand(cmd.ModelCmd)
	cobra.AddCommand(cmd.RequestCmd)
	cobra.AddCommand(cmd.ResponseCmd)
	cobra.AddCommand(cmd.ServiceCmd)
	cobra.AddCommand(cmd.EnumCmd)
}
