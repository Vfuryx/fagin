package cli

//import (
//	"fadmin/pkg/controller"
//	"fadmin/utils"
//	"fmt"
//	"path"
//
//	"github.com/spf13/cobra"
//)
//
//var ControllerCmd = &cobra.Command{
//	Use:   "controller",
//	Short: "控制器",
//	Long: `
//该子命令支持生成控制器模版，用法如下：
//go run console/main.go controller path
//`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("controller called")
//		if len(args) == 0 {
//			fmt.Println("create controller err")
//			return
//		}
//		p := utils.PathUnderscore(args[0])
//		controller.CreateControllerTemplate(p, path.Base(p))
//	},
//}
