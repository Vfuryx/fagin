package consoles

import (
	"fagin/config"
	"fagin/console"
	"fagin/pkg/server"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "",
	Short: "开启服务",
	Long:  `开启服务`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

// Execute 执行
func Execute() {
	config.Init()
	console.Handle(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
