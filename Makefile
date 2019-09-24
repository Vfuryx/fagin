export GOPROXY=https://goproxy.cn
export GO111MODULE=on

Name=$(name)

default:
	go run main.go

install:
	go mod tidy
	go mod vendor

doc:
	swag init

create-cmd:
	cd ./console && cobra add ${name}

create-admin:
	go run console/main.go admin

migrate:
	go run console/main.go migrate

migrate\:create:
	go run console/main.go migrate create ${Name}

migrate\:reset:
	go run console/main.go migrate reset

migrate\:rollback:
	go run console/main.go migrate rollback

build:
	go build main.go

jwt-secret:
	echo "jwt: \n  secret: `head -c 32 /dev/random | base64`" >> .env.yml


help:
	@echo "make - 运行服务"
	@echo "make install - 更新和载入包文件"
	@echo "make doc - 生成api文档"
	@echo "make create-cmd - 创建一个 cli 模版"
	@echo "make create-admin - 创建一个基本的权限结构"
	@echo "make migrate - 迁移数据库"
	@echo "migrate\:reset - 数据库迁移重置"
	@echo "migrate\:rollback - 数据库迁移回滚"
	@echo "migrate\:build - 打包"
	@echo "migrate jwt-secret - 生成 jwt 密钥"