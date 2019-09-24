# fagin

gin 构建的 api 项目

扩展包
| 包 | 详情 |
| ---|--- |
| gin | gin |
| gorm | orm |
| gormigrate | 数据迁移 |
| spf13/viper | 配置 |
| spf13/cobra  | cli |
| swaggo/gin-swagger | 自动生成RESTful API文档的gin中间件 |
| swaggo/gin-swagger/swaggerFiles | swagger嵌入文件 | 
| casbin/casbin | 权限配置 |
| casbin/gorm-adapter | Gorm适配器。使用此库，Casbin可 以从Gorm支持的数据库加载策略或将策略保存到它。 |
| sirupsen/logrus  | 日志包 |
| lestrrat/go-file-rotatelogs | 日志文件分割 |

## 项目目录结构
<details>
<summary>展开查看</summary>
<pre><code>
├── app                 项目核心逻辑代码
│    ├── constants      常量
│    ├── errno          错误
│    ├── controllers    控制器
│    ├── models         模型
│    ├── middleware     中间件
│    ├── services       业务
│    ├── helper.go      工具方法
│    └── requests       参数校验
│
├── console         控制台
│    ├── cmd        cmd 入口
│    └── routes.go  cmd 命令
│
├── config           配置中心
│
├── database         数据库
│    └── migrations  数据迁移
│
├── pkg              项目依赖
│
├── doc              文档
│
├── public           项目静态文件
│    ├── css     
│    ├── js      
│    ├── js      
│    └── assets  
│
├── resources        前端源码
│    ├── assets        assets
│    └── view        go 模板文件
│
├── routes           路由
│    ├── api.go      api 路由注册
│    └── web.go      页面路由注册
│
├── storage          存放日志等文件
│
├── main.go          项目入口
│
├── config.yaml      项目配置
│
├── .env             系统配置
│
├── .drone.yml       drone 配置
│
├── admin.sh         脚本
│
└── Makefile         Makefile 文件
</code></pre>
</details>

## 启动web server

```
 make
```

## 安装项目依赖库

```
 make install 
```

## 数据迁移

```
 make migrate
```

## 创建数据迁移

```
 make migrate:create name=create_tablename
```

## 数据迁移重置

```
 make migrate:reset
```

## 数据迁移回滚一步

```
 make migrate:rollback
```

## 打包程序

```
 make build
```

## 生成 api 文档 查看文档 

http://localhost:8080/swagger/index.html

```
 make doc 
```

## 生成jwt的secret 

```
 make jwt-secret
```

## 生成 CMD 命令 「 name 为命令名称」

```
 make create-cmd name=？
```

## 生成一个超级管理员账户 admin 123456

```
 make create-admin
```


> 提示 如果无法下载 golang.org/x 等包，可用 proxy 的方式下载包

```
export GOPROXY=https://goproxy.cn
```




### CICD方式

1. gitlab + jenkins

1. gitea + drone

