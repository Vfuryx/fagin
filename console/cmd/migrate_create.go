package cmd

import (
	"fagin/pkg/migrate"
	"fmt"
	"github.com/spf13/cobra"
)

// MigrateCreateCmd 创建一个新的数据迁移
var MigrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建一个新的数据迁移",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("文件名不能为空")
			return
		}

		err := migrate.CreateMigration(args[0])
		if err != nil {
			panic(err)
		}
	},
}
