package cli

import (
	"fagin/pkg/enum"
	"fagin/utils"
	"fmt"
	"path"

	"github.com/spf13/cobra"
)

var EnumCmd = &cobra.Command{
	Use:   "enum",
	Short: "枚举类型",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enum called")
		if len(args) == 0 {
			fmt.Println("create enum err")
			return
		}
		p := utils.PathUnderscore(args[0])
		enum.CreateEnumTemplate(p, path.Base(p))
	},
}
