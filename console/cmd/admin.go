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
	"fagin/app/models/role"
	"fagin/app/models/user"
	"fagin/app/service"
	"fagin/pkg/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("admin called")

		password, _ := service.Encrypt("123456")
		u := user.User{
			Username: "admin",
			Email:    "admin@admin.com",
			Password: password,
		}

		// 添加超级管理员账号
		if ok, _ := user.Create(&u); !ok {
			fmt.Println("创建超级管理员失败")
			return
		}
		// 添加角色
		r := role.Role{Name: "admin", Sort: 9999}
		role.Create(&r)

		// 设置角色权限
		service.Canbin.AddPolicyForRole("admin", "/*", "|")

		// 设置超级管理员角色
		service.Canbin.AddUserRole(strconv.Itoa(int(u.ID)), r.Name)

		db.ORM.Close()
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
