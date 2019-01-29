# fagin

gin 构建的 api 项目

扩展包

- gin

- gorm

- github.com/spf13/viper@v1.3.1


### 项目结构

```
- app           -- app
-- constants    -- 常量
-- models       -- 模型
-- controllers  -- 控制器
-- middleware   -- 中间件
-- services     -- 业务
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


- make install 

安装项目依赖库

- make dev 

启动web server
