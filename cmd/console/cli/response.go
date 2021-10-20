package cli

import (
	"fagin/pkg/response"
	"fagin/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var ResponseCmd = &cobra.Command{
	Use:   "response",
	Short: "响应",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("response called")
		if len(args) <= 0 {
			fmt.Println("create response err")
			return
		}
		path, name := utils.GetPathAndUnderscore(args[0])
		response.CreateResponseTemplate(path, name)
	},
}
