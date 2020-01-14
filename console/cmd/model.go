package cmd

import (
	"fagin/pkg/db"
	"fmt"
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
		if len(args) <= 0 {
			fmt.Println("create model err")
			return
		}
		path, name := Handle(args[0])
		db.CreateModelTemplate(path, name)
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
