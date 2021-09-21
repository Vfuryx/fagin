package cmd

import (
	"fagin/pkg/enum"
	"fagin/pkg/paths"
	"fmt"
	"github.com/spf13/cobra"
)

var EnumCmd = &cobra.Command{
	Use:   "enum",
	Short: "枚举类型",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enum called")
		if len(args) <= 0 {
			fmt.Println("create enum err")
			return
		}
		path, name := paths.GetPathAndUnderscore(args[0])
		enum.CreateEnumTemplate(path, name)
	},
}
