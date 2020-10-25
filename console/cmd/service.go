package cmd

import (
	"fagin/pkg/service"
	"fmt"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
		if len(args) <= 0 {
			fmt.Println("create service err")
			return
		}
		path, name := Handle(args[0])
		service.CreateServiceTemplate(path, name)
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
