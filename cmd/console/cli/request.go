package cli

//import (
//	"fadmin/pkg/request"
//	"fadmin/utils"
//	"fmt"
//	"path"
//
//	"github.com/spf13/cobra"
//)
//
//var RequestCmd = &cobra.Command{
//	Use:   "request",
//	Short: "表单验证",
//	Long:  ``,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("request called")
//		if len(args) == 0 {
//			fmt.Println("create request err")
//			return
//		}
//		p := utils.PathUnderscore(args[0])
//		request.CreateRequestTemplate(p, path.Base(p))
//	},
//}
