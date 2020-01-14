// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"unicode"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
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