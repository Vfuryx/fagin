export GOPROXY=https://goproxy.cn
export GO111MODULE=on

Name=$(name)

default:
	air -c ./.air.toml

install:
	go mod tidy
	go mod vendor

web:
	cd ./resources/assets/vben-admin-thin-next && yarn run dev

web\:build:
	cd ./resources/assets/vben-admin-thin-next && yarn run build

doc:
	#需要权限来修改文件
	sudo swag init -o ./docs/swag

create-cmd:
	cd ./console && cobra add ${name}

h:
	go run main.go help

enum:
	go run main.go enum ${Name}

request:
	go run main.go request ${Name}

response:
	go run main.go response ${Name}

model:
	go run main.go model ${Name}

cache:
	go run main.go cache ${Name}

controller:
	go run main.go controller ${Name}

service:
	go run main.go service ${Name}

create-admin:
	go run main.go admin

migrate:
	go run main.go migrate

migrate\:create:
	go run main.go migrate create ${Name}

migrate\:reset:
	go run main.go migrate reset

migrate\:rollback:
	go run main.go migrate rollback

build:
	go build -o ${Name} .

build\:winamd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${Name} .

build\:win386:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ${Name} .

build\:darwin386:
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ${Name} .

build\:darwinamd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${Name} .

build\:darwinarm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm go build -o ${Name} .

build\:darwinarm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${Name} .

build\:linux386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ${Name} .

build\:linuxamd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${Name} .

build\:linuxarm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ${Name} .

build\:linuxarm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${Name} .

key\:generate:
	go run console/main.go key

jwt-secret:
	echo "jwt: \n  secret: `head -c 32 /dev/random | base64`" >> .env.yml

pipeline:
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod tidy
	go run main.go migrate
	go build -o ${Name} .
	sh admin.sh restart ${Name}
	sh admin.sh status ${Name}

start:
	sh admin.sh restart ${Name}

stop:
	sh admin.sh stop ${Name}

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
	@echo "make enum            	- 生成 枚举"
	@echo "make request             - 生成 请求验证"
	@echo "make response            - 生成 生成响应"
	@echo "make model               - 生成 模型"
	@echo "make cache               - 生成 缓存"
	@echo "make controller          - 生成 控制器"
	@echo "make service             - 生成 服务"
	@echo "make pipeline            - 流水线"
	@echo "make start            	- 运行服务"
	@echo "make stop            	- 关闭服务"