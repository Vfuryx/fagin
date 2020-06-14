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
	Short: "生成一个超级管理员账户",
	Long: `详细说明`,
	Example: "例子",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("admin called")

		u := user.User{
			Username: "admin",
			Email:    "admin@admin.com",
			Password: "123456",
		}

		// 添加超级管理员账号
		err := user.Dao().Create(&u)
		if err != nil {
			fmt.Println("创建超级管理员失败")
			return
		}
		// 添加角色
		r := role.Role{Name: "admin", Sort: 9999}
		_ = role.Dao().Create(&r)

		// 设置角色权限
		service.Canbin.AddPolicyForRole("admin", "/*", "|")

		// 设置超级管理员角色
		_, _ = service.Canbin.AddUserRole(strconv.Itoa(int(u.ID)), r.Name)

		db.ORM.Close()
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
