package cmd

import (
	"fagin/pkg/cache"
	"fagin/pkg/paths"
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
		path, name := paths.GetPathAndUnderscore(args[0])
		cache.CreateCacheTemplate(path, name)
	},
}
