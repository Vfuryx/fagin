package cli

//import (
//	"fadmin/pkg/service"
//	"fadmin/utils"
//	"fmt"
//	"path"
//
//	"github.com/spf13/cobra"
//)
//
//var ServiceCmd = &cobra.Command{
//	Use:   "service",
//	Short: "生成服务模版",
//	Long: `
//该子命令支持生成服务模版，用法如下：
//go run console/main.go service path
//`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("service called")
//		if len(args) == 0 {
//			fmt.Println("create service err")
//			return
//		}
//		p := utils.PathUnderscore(args[0])
//		service.CreateServiceTemplate(p, path.Base(p))
//	},
//}
