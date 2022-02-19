package cli

import (
	"fagin/pkg/cache"
	"fagin/utils"
	"fmt"
	"path"

	"github.com/spf13/cobra"
)

var CacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "缓存",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cache called")
		if len(args) == 0 {
			fmt.Println("create cache err")
			return
		}

		p := utils.PathUnderscore(args[0])
		cache.CreateCacheTemplate(p, path.Base(p))
	},
}
