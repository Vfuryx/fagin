package cmd

import (
	"fagin/pkg/paths"
	"fagin/pkg/request"
	"fmt"
	"github.com/spf13/cobra"
)

var RequestCmd = &cobra.Command{
	Use:   "request",
	Short: "表单验证",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("request called")
		if len(args) <= 0 {
			fmt.Println("create request err")
			return
		}
		path, name := paths.GetPathAndUnderscore(args[0])
		request.CreateRequestTemplate(path, name)
	},
}
