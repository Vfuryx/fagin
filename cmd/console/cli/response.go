package cli

import (
	"fadmin/pkg/response"
	"fadmin/utils"
	"fmt"
	"path"

	"github.com/spf13/cobra"
)

// ResponseCmd 响应
var ResponseCmd = &cobra.Command{
	Use:   "response",
	Short: "响应",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("response called")
		if len(args) == 0 {
			fmt.Println("create response err")
			return
		}

		p := utils.PathUnderscore(args[0])
		response.CreateResponseTemplate(p, path.Base(p))
	},
}
