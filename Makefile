
default:
	go run main.go

install:
	govendor sync -v

migrate:
	go run console/main.go migrate

build:
	go build .