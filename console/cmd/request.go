package cmd

import (
	"fagin/pkg/request"
	"fmt"
	"github.com/spf13/cobra"
)

// requestCmd represents the request command
var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("request called")
		if len(args) <= 0 {
			fmt.Println("create request err")
			return
		}
		path, name := Handle(args[0])
		request.CreateRequestTemplate(path, name)
	},
}

func init() {
	rootCmd.AddCommand(requestCmd)
}