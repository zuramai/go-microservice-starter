install:
	go mod download
	go mod vendor

runserver:
	go run cmd/grpc/main.go

runtest:
	go clean -testcache
	gotest ./test -v

generateproto:
	bash /applicationservice/client/generate.sh

docker-push:
	@docker push

stop:
	@docker-compose down

.PHONY: install runtest generateproto runserver