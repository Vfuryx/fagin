# fagin

gin 构建的 api 项目

扩展包

- gin                                           ---- gin
- gorm                                          ---- orm
- gormigrate                                    ---- 数据迁移
- github.com/spf13/viper@v1.3.1                 ---- 配置
- github.com/spf13/cobra                        ---- cli 
- github.com/casbin/casbin                      ---- 权限配置
- github.com/swaggo/gin-swagger                 ---- 自动生成RESTful API文档的gin中间件
- github.com/swaggo/gin-swagger/swaggerFiles    ---- swagger嵌入文件
- github.com/casbin/gorm-adapter                ---- Gorm适配器。使用此库，Casbin可以从Gorm支持的数据库加载策略或将策略保存到它。
- github.com/sirupsen/logrus                    ---- 日志包
- github.com/lestrrat/go-file-rotatelogs        ---- 日志文件分割

### 项目结构

```
- app           -- app
-- constants    -- 常量
-- models       -- 模型
-- controllers  -- 控制器
-- middleware   -- 中间件
-- services     -- 业务
- console       -- 控制台
-- main.go      -- cmd 入口
-- cmd          -- cmd 命令
- config        -- 项目配置
- database      -- 数据库
- public        -- 公共口
-- css
-- js
-- img
-- assets       -- assets
- routes        -- 路由
- tests         -- 测试
- resources     -- 资源
-- assets       -- assets
-- views        -- 模版
- storage       -- 存储位置
-- logs         -- 错误日志
- pak           -- 框架

- main.go       -- 入口文件
- .env.yml      -- 系统配置
- config.yml    -- 项目配置
- .drone.yml    -- drone 配置
- Makefile      -- makefile


```

- make 

    启动web server

- make install 

    安装项目依赖库

- make migrate

    数据迁移

- make migrate:reset

    数据迁移重置

- make migrate:rollback

    数据迁移回滚一步

- make build

    打包程序
    
- make doc 

    生成 api 文档 查看文档 http://localhost:8080/swagger/index.html

- make jwt-secret
    
    生成jwt的secret
    
- make create-cmd name=？
    
    生成 CMD 命令 「 name 为命令名称」

- make create-admin
    
    生成一个超级管理员账户 admin 123456



> 提示 如果无法下载 golang.org/x 等包，可用 proxy 的方式下载包

```
export GOPROXY=https://goproxy.io
```




### CICD方式

1. gitlab + jenkins

1. gitea + drone

