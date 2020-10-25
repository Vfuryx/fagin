package cmd

import (
	"fagin/pkg/controller"
	"fmt"

	"github.com/spf13/cobra"
)


var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("controller called")
		if len(args) <= 0 {
			fmt.Println("create controller err")
			return
		}
		path, name := Handle(args[0])
		controller.CreateControllerTemplate(path, name)
	},
}

func init() {
	rootCmd.AddCommand(controllerCmd)
}
