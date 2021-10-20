# fagin

Gin 构建的 web 项目

## 扩展包

 包 | 详情 
 ---|--- 
 gin | Gin 
 gorm V2 | ORM 
 sql-migrate | 数据迁移 
 spf13/viper | 配置 
 spf13/cobra  | CLI
 swaggo | Swagger 
 casbin V2 | 权限配置 
 logrus  | 日志包 
 go-file-rotatelogs | 日志文件分割 
 gin-contrib/sessions | Session包
 go-redis V8 | 缓存
 BigCache | 缓存
 tableflip | 进程重启

# 热更新 air
热更新需要安装 air

## 项目目录结构
<details>
<summary>展开查看</summary>
<pre><code>
├── app                 项目核心逻辑代码
│    ├── caches         缓存
│    ├── enums          枚举
│    ├── errno          错误码
│    ├── controllers    控制器
│    ├── models         模型
│    ├── middleware     中间件
│    ├── service        服务
│    ├── requests       参数校验
│    ├── responses      响应
│    └── helper.go      工具方法
│
├── console             控制台
│    ├── cmd            cmd 文件
│    └── console.go     cmd 加载
│
├── config              配置中心
│
├── database            数据库
│    └── migrations     数据迁移文件
│
├── pkg                 项目依赖
│
├── doc                 文档
│
├── public              项目公开静态文件（包含上传文件）
│    ├── css     
│    ├── js      
│    ├── js      
│    ├── upload      
│    └── assets
|
├── resources           项目资源
│    ├── assets         assets
│    │    └── admin     后台前端项目
│    ├── static         打包静态文件
│    └── views          模板文件
│
├── routes              路由
│    ├── root.go        路由加载
│    ├── admin.go       后台 路由
│    ├── api.go         api 路由
│    └── web.go         web 路由
│ 			
├── utils               工具
├── test                测试文件夹
├── storage             存放日志等文件
├── .air.toml           热更新配置
├── main.go             项目入口
├── config.yaml         动态项目配置
├── .drone.yml          drone 配置
├── admin.sh            脚本
└── Makefile            Makefile 文件
</code></pre>
</details>


### 启动web server

```
 make
```

#### 安装项目依赖库

```
 make install 
```

#### 查看命令行帮助

```
 make h
```
#### 生成表枚举模版 「 name 为文件路径加名称」

```
 make enum name=TestType // 或者 test_type
```

#### 生成表单验证模版 「 name 为文件路径加名称」

```
 make request name=/Api/TestRequest // 或者 test_request
```

#### 生成响应模版 「 name 为文件路径加名称」

```
 make response name=/Api/testResponse	// 或者 test_response
```

#### 生成模型模版 「 name 为文件路径加名称」

```
 make model name=testModel	// 或者 test_model
```

#### 生成缓存模版 「 name 为文件路径加名称」

```
 make cache name=testCache	// 或者 test_cache
```

#### 生成控制器模版 「 name 为文件路径加名称」

```
 make controller name=admin/testController	// 或者 admin/test_controller
```

#### 生成服务模版 「 name 为文件路径加名称」

```
 make service name=test	// 或者 admin/test
```

#### 数据迁移

```
 make migrate
```

#### 创建数据迁移

```
 make migrate:create name=create_tablename
```

#### 数据迁移重置

```
 make migrate:reset
```

#### 数据迁移回滚一步

```
 make migrate:rollback
```

#### 打包程序

```
 make build [name=xxx] 
```
#### 运行程序

```
 make start name=xxx
```
#### 关闭程序

```
 make stop name=xxx
```

#### 流水线

```
 make pipeline name=xxx
```

#### 交叉编译
<details>
<summary>展开查看</summary>
<pre><code>
生成 win64位程序
make build:winamd64 [name=xxx]

生成 win32位程序
make build:win386 [name=xxx]

生成 darwin386 程序
make build:darwin386 [name=xxx]

生成 darwinamd64 程序
make build:darwinamd64 [name=xxx]

生成 darwinarm程序
make build:darwinarm [name=xxx]

生成 darwinarm64程序
make build:darwinarm64 [name=xxx]

生成 linux386程序
make build:linux386 [name=xxx]

生成 linuxamd64程序
make build:linuxamd64 [name=xxx]

生成 linuxarm程序
make build:linuxarm [name=xxx]

生成 linuxarm64程序
make build:linuxarm64 [name=xxx]


</code></pre>
</details>

#### 生成 api 文档 查看文档 

http://localhost:8080/swagger/index.html

```
 make doc 
```

#### 生成应用的 key（会覆盖）

```
 make key:generate 
```

#### 生成jwt的secret 

```
 make jwt-secret
```

#### 生成 CMD 命令 「 name 为命令名称」

```
 make create-cmd name=？
```

#### 生成一个超级管理员账户 admin 123456

```
 make create-admin
```

#### session 使用

```
session := sessions.Default(ctx)
if id := session.Get("user"); id == nil {
	session.Set("user", 100)
	// 将 sessionID 存入cookie
	if err := session.Save(); err != nil {
		log.Info(err)
	}
}
```


---

> 提示 如果无法下载 golang.org/x 等包，可用 proxy 的方式下载包

```
export GOPROXY=https://goproxy.cn,direct
# 或者
go env -w GOPROXY=https://goproxy.cn,direct
```




### CICD方式

1. gitlab + jenkins

1. gitea + drone

## 总后台 前端

#### 快速运行
```
make web
```
#### 快速打包
```
make web:build
```
#### 位置
```
cd /resources/assets/admin 
```
#### install npm
```
npm i 
```
#### 运行项目
```
npm run dev
```
#### 新增并设置生产环境配置 （参考 .env.development 配置）
```
vim .env.production
```
#### 打包项目

```
npm run build:prod  (打包生产环境配置)
// 打包的文件会生成在指定位置，无需更改
```

## 开启后端服务

```
# 新建 env
cp .config.yml.example .config.yml

# 配置 env
vim .config.yml

# 运行
make 
or
go run .

# 打包项目
make build name=main
or
go build .

```


## 进程守护 systemd （ 适配 tableflip ） 

> tableflip 仅适用于 Linux 和 macOS。
> 运行程序后，会在根目录下生成 pid 文件。

配合 `make pipeline name={name}` 命令，实现 零停机 平滑重启 。

```
[Unit]
Description=Service

[Service]
ExecStart=/path/to/binary -some-flag /path/to/pid-file
ExecReload=/bin/kill -HUP $MAINPID
PIDFile=/path/to/pid-file
```

示例

```
# /etc/systemd/system/app.service

[Unit]
Description=Service

[Service]
ExecStart=/work_path/main
ExecReload=/bin/kill -HUP $MAINPID
PIDFile=/work_path/app.pid
```