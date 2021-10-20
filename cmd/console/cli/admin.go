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

package cli

import (
	"fagin/app"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	"fagin/pkg/casbins"
	"fagin/pkg/db"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// AdminCmd represents the admin command
var AdminCmd = &cobra.Command{
	Use:     "admin",
	Short:   "生成一个超级管理员账户",
	Long:    `详细说明`,
	Example: "例子",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("admin called")
		pwd, _ := app.Encrypt("12345678")
		u := admin_user.AdminUser{
			ID:       1,
			Username: "admin",
			NickName: "admin",
			Password: pwd,
			Avatar:   "http://qiniu.furyx.cn/photo.jpg",
			Status:   1,
		}

		// 添加超级管理员账号
		err := db.ORM().FirstOrCreate(&u).Error
		if err != nil {
			fmt.Println("创建超级管理员失败")
			return
		}
		// 添加角色
		r := admin_role.AdminRole{
			ID:     1,
			Name:   "admin",
			Key:    "admin",
			Remark: "admin",
		}
		err = db.ORM().FirstOrCreate(&r).Error
		if err != nil {
			fmt.Println("创建角色失败")
			return
		}

		// 设置角色权限
		ok, err := casbins.Casbin.AddPolicyForRole("admin", "/*", "|")
		if err != nil || !ok {
			fmt.Println("设置角色权限失败", err)
			return
		}

		// 设置超级管理员角色
		ok, err = casbins.Casbin.AddUserRole(strconv.Itoa(int(u.ID)), r.Name)
		if err != nil || !ok {
			fmt.Println("设置超级管理员角色失败", err)
			return
		}

		//db.ORM().Close()
	},
}
