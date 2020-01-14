package cmd

import (
	"fagin/pkg/response"
	"fmt"
	"github.com/spf13/cobra"
)

// responseCmd represents the response command
var responseCmd = &cobra.Command{
	Use:   "response",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("response called")
		if len(args) <= 0 {
			fmt.Println("create response err")
			return
		}
		path, name := Handle(args[0])
		response.CreateResponseTemplate(path, name)
	},
}

func init() {
	rootCmd.AddCommand(responseCmd)
}