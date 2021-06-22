export GOPROXY=https://goproxy.cn
export GO111MODULE=on

Name=$(name)

default:
	air -c ./.air.toml

install:
	go mod tidy
	go mod vendor

web:
	cd ./resources/assets/admin2 && yarn run serve

web\:build:
	cd ./resources/assets/admin2 && yarn run build

doc:
#需要权限来修改文件
	sudo swag init

create-cmd:
	cd ./console && cobra add ${name}

request:
	go run main.go -c request ${Name}

response:
	go run main.go -c response ${Name}

model:
	go run main.go -c model ${Name}

cache:
	go run main.go -c cache ${Name}

controller:
	go run main.go -c controller ${Name}

service:
	go run main.go -c service ${Name}

create-admin:
	go run main.go -c admin

migrate:
	go run main.go -c migrate

migrate\:create:
	go run main.go -c migrate create ${Name}

migrate\:reset:
	go run main.go -c migrate reset

migrate\:rollback:
	go run main.go -c migrate rollback

build:
	go build main.go

build\:winamd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build .

build\:win386:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build .

build\:darwin386:
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build .

build\:darwinamd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build .

build\:darwinarm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm go build .

build\:darwinarm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build .

build\:linux386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build .

build\:linuxamd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

build\:linuxarm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build .

build\:linuxarm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build .

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
	@echo "make build:winamd64      - 打包程序winamd64"
	@echo "make build:win386      	- 打包程序win386"
	@echo "make build:darwin386     - 打包程序darwin386"
	@echo "make build:darwinamd64   - 打包程序darwinamd64"
	@echo "make build:darwinarm     - 打包程序darwinarm"
	@echo "make build:darwinarm64   - 打包程序darwinarm64"
	@echo "make build:linux386      - 打包程序linux386"
	@echo "make build:linuxamd64    - 打包程序linuxamd64"
	@echo "make build:linuxarm      - 打包程序linuxarm"
	@echo "make build:linuxarm64    - 打包程序linuxarm64"
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
	@echo "make cache               - 生成 缓存"
	@echo "make controller          - 生成 控制器"
	@echo "make service             - 生成 服务"
	@echo "make pipeline            - 流水线"