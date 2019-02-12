
default:
	go run main.go

install:
	govendor sync -v

migrate:
	go run console/main.go migrate

migrate\:reset:
	go run console/main.go migrate reset

migrate\:rollback:
	go run console/main.go migrate rollback

build:
	go build .