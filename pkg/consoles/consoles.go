package consoles

import (
	"fagin/console"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "console",
	Short: "",
	Long:  ``,
}

var IsConsole bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&IsConsole, "console", "c", false, "-c 代表命令模式")
}

func Execute() {
	console.Handle(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
