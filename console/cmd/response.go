package cmd

import (
	"fagin/pkg/paths"
	"fagin/pkg/response"
	"fmt"
	"github.com/spf13/cobra"
)

var ResponseCmd = &cobra.Command{
	Use:   "response",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("response called")
		if len(args) <= 0 {
			fmt.Println("create response err")
			return
		}
		path, name := paths.GetPathAndUnderscore(args[0])
		response.CreateResponseTemplate(path, name)
	},
}

