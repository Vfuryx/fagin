package cli

import (
	"fagin/pkg/cache"
	"fagin/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var CacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "缓存",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cache called")
		if len(args) <= 0 {
			fmt.Println("create cache err")
			return
		}
		path, name := utils.GetPathAndUnderscore(args[0])
		cache.CreateCacheTemplate(path, name)
	},
}
