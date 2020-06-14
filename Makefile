export GOPROXY=https://goproxy.cn
export GO111MODULE=on

Name=$(name)

default:
	go run main.go

install:
	go mod tidy
	go mod vendor

doc:
#需要权限来修改文件
	sudo swag init

create-cmd:
	cd ./console && cobra add ${name}

request:
	go run console/main.go request ${Name}

response:
	go run console/main.go response ${Name}

model:
	go run console/main.go model ${Name}

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

key\:generate:
	go run console/main.go key

jwt-secret:
	echo "jwt: \n  secret: `head -c 32 /dev/random | base64`" >> .env.yml

pipeline:
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod tidy
	go run console/main.go migrate
	go build -o fagin ./main.go
	sh admin.sh restart fagin
	sh ./admin.sh status fagin

help:
	@echo "make                     - 运行服务"
	@echo "make build               - 打包程序"
	@echo "make install             - 更新和载入包文件"
	@echo "make doc                 - 生成api文档"
	@echo "make create-cmd          - 创建一个 cli 模版"
	@echo "make create-admin        - 创建一个基本的权限结构"
	@echo "make migrate             - 迁移数据库"
	@echo "make migrate:reset       - 数据库迁移重置"
	@echo "make migrate:rollback    - 数据库迁移回滚"
	@echo "make jwt-secret          - 生成 jwt 密钥"
	@echo "make key:generate        - 生成 key"
	@echo "make request             - 生成 请求验证"
	@echo "make response            - 生成 生成响应"
	@echo "make model               - 生成 模型"
	@echo "make pipeline            - 流水线"