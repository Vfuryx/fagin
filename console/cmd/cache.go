package cmd

import (
	"fagin/pkg/cache"
	"fmt"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cache called")
		if len(args) <= 0 {
			fmt.Println("create cache err")
			return
		}
		cache.CreateCacheTemplate(args[0])
	},
}

func init() {
	rootCmd.AddCommand(cacheCmd)
}
