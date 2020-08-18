# fagin

gin 构建的 api 项目

## 扩展包

 包 | 详情 
 ---|--- 
 gin | gin 
 gorm | orm 
 sql-migrate | 数据迁移 
 spf13/viper | 配置 
 spf13/cobra  | cli 
 swaggo/gin-swagger | 自动生成RESTful API文档的gin中间件 
 swaggo/gin-swagger/swaggerFiles | swagger嵌入文件 |
 casbin/casbin | 权限配置 
 casbin/gorm-adapter | Gorm适配器。使用此库，Casbin可 以从Gorm支持的数据库加载策略或将策略保存到它。 
 sirupsen/logrus  | 日志包 
 lestrrat/go-file-rotatelogs | 日志文件分割 
 gin-contrib/sessions | session包
 go-redis/redis/v7 | 缓存

## 项目目录结构
<details>
<summary>展开查看</summary>
<pre><code>
├── app                 项目核心逻辑代码
│    ├── cache          缓存
│    ├── constants      常量
│    ├── errno          错误
│    ├── controllers    控制器
│    ├── models         模型
│    ├── middleware     中间件
│    ├── services       业务
│    ├── requests       参数校验
│    ├── responses      响应
│    └── helper.go      工具方法
│
├── console             控制台
│    ├── cmd            cmd 入口
│    └── main.go        cmd 命令
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
├── public              项目静态文件
│    ├── css     
│    ├── js      
│    ├── js      
│    └── assets  
│
├── resources           项目资源
│    ├── assets         assets
│    │     └── admin    后台前端项目
│    └── views          go 模板文件
│
├── routes              路由
│    ├── admin.go       后台 路由
│    ├── api.go         api 路由
│    └── web.go         web 路由
│ 			
├── test                测试文件  
│
├── storage             存放日志等文件
│
├── main.go             项目入口
│
├── config.yaml         动态项目配置
│
├── .env                系统配置
│
├── .drone.yml          drone 配置
│
├── admin.sh            脚本
│
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
 make build
```

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
export GOPROXY=https://goproxy.cn
```




### CICD方式

1. gitlab + jenkins

1. gitea + drone

## admin 前端

```
# 位置 
cd /resources/assets/admin 

# install npm
npm i 

# 运行项目
npm run dev

#新增并设置生产环境配置 （参考 .env.development 配置）
vim .env.production

# 打包项目
npm run build:prod  (打包生产环境配置)

// 打包的文件会生成在指定位置，无需更改

```

## 开启后端服务

```
# 新建 env
cp .env.example .env

# 配置 env
vim .env

# 运行
make 
or
go run .

# 打包项目
make build
or
go build .

```
