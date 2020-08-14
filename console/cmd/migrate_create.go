package cmd

import (
	"fagin/pkg/db"
	"fagin/pkg/migrate"
	"fmt"
	"github.com/spf13/cobra"
)

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建一个新的数据迁移",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("文件名不能为空")
			return
		}
		migrate.Create(args[0])
		// 关闭orm
		defer db.Close()
	},
}

func init() {
	// migrate 的子命令
	migrateCmd.AddCommand(migrateCreateCmd)
}
