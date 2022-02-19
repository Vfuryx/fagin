package cli

import (
	"fagin/pkg/db"
	"fagin/utils"
	"fmt"
	"path"

	"github.com/spf13/cobra"
)

var ModelCmd = &cobra.Command{
	Use:   "model",
	Short: "模型",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
		if len(args) == 0 {
			fmt.Println("create model err")
			return
		}
		p := utils.PathUnderscore(args[0])
		db.CreateModelTemplate(p, path.Base(p))
	},
}
