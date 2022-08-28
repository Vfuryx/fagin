package cli

import (
	_ "fagin/database/migrations"
	"fagin/pkg/migrate"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "数据迁移",
	Long:    ``,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Migration()
		log.Printf("Migration did run successfully")
	},
}

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

// MigrateResetCmd 重置数据迁移
var MigrateResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "重置数据迁移",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Reset()
		log.Printf("Migration did run successfully")

	},
}

// MigrateRollbackCmd 回滚到上一步
var MigrateRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "回滚到上一步",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Rollback()
		log.Printf("Migration did run successfully")

	},
}

// MigrateRefreshCmd 回滚所有迁移，并运行所有迁移
var MigrateRefreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "回滚到上一步",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Refresh()
		log.Printf("Migration did run successfully")

	},
}
