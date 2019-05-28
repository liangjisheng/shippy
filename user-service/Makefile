build:
	# protoc -I. --go_out=.:$(GOPATH)/src/shippy/user-service --micro_out=.:$(GOPATH)/src/shippy/user-service proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o user-service
	docker build -t user-service .
	rm user-service
run:
	# 使用环境变量来决定是否调用其他服务，从而使你的微服务能隔离的进行测试
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns -e DISABLE_AUTH=true user-service