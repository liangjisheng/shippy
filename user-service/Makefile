build:
	# protoc -I. --go_out=.:$(GOPATH)/src/shippy/user-service --micro_out=.:$(GOPATH)/src/shippy/user-service proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o user-service
	docker build -t user-service .
	# rm user-service
run:
	# 使用环境变量来决定是否调用其他服务，从而使你的微服务能隔离的进行测试
	# docker run --net="host"\
	# 	-e MICRO_REGISTRY=mdns \
	# 	-e DISABLE_AUTH=true \
	# 	user-service

	docker run --net="host" -p 5054 \
		-e MICRO_REGISTRY=mdns \
		-e DB_NAME="postgres" \
		-e DB_HOST="127.0.0.1" \
		-e DB_PORT="5432" \
		-e DB_USER="postgres" \
		-e DB_PASSWORD="postgres" \
		-e MICRO_BROKER=nats \
		-e MICRO_BROKER_ADDRESS=0.0.0.0:4222 \
		user-service