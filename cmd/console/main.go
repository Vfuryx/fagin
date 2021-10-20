package main

import (
	"fagin/cmd/console/cli"
	"fagin/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(cli.AdminCmd)
	rootCmd.AddCommand(cli.CacheCmd)
	rootCmd.AddCommand(cli.ControllerCmd)
	rootCmd.AddCommand(cli.KeyCmd)
	cli.MigrateCmd.AddCommand(cli.MigrateCreateCmd)
	cli.MigrateCmd.AddCommand(cli.MigrateResetCmd)
	cli.MigrateCmd.AddCommand(cli.MigrateRollbackCmd)
	rootCmd.AddCommand(cli.MigrateCmd)
	rootCmd.AddCommand(cli.ModelCmd)
	rootCmd.AddCommand(cli.RequestCmd)
	rootCmd.AddCommand(cli.ResponseCmd)
	rootCmd.AddCommand(cli.ServiceCmd)
	rootCmd.AddCommand(cli.EnumCmd)
}

func main() {
	// 初始化配置
	config.Init()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
