package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"unicode"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "",
	Long:  ``,
}


func Execute() {

	// 生成文档
	//err := doc.GenMarkdownTree(rootCmd, "./")
	//if err != nil {
	//	fmt.Println(err)
	//}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 驼峰式写法转为下划线写法
//将大写转成 _a 的形式
func Underscore(str string) string {
	var buf strings.Builder
	for index, value := range str {
		// value >= 'A' && value <= 'Z'
		if unicode.IsUpper(value) {
			if index != 0 {
				buf.WriteByte('_')
			}
			// byte(value + 'a' - 'A')
			buf.WriteRune(unicode.ToLower(value))
			continue
		}
		buf.WriteRune(value)
	}
	return buf.String()
}

// 下划线写法转为驼峰写法
func Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func Handle(name string) (string, string) {
	str := strings.Trim(name, " /\\")
	sl := strings.Split(str, "/")
	for index, value := range sl {
		sl[index] = Underscore(value)
	}
	path := strings.Join(sl, "/")
	return path, sl[len(sl)-1]
}