# fagin

gin 构建的 api 项目

扩展包

- gin                                   ---- gin
- gorm                                  ---- orm
- gormigrate                            ---- 数据迁移
- github.com/spf13/viper@v1.3.1         ---- 配置
- github.com/spf13/cobra                ---- cli 

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

- main.go       -- 入口文件
- env.yml       -- 配置文件
- Makefile      -- makefile


```

- make 

    启动web server

- make install 

    安装项目依赖库

- make migrate

    数据迁移


