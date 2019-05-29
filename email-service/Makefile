build:
	GOOS=linux GOARCH=amd64 go build -o email-service
	docker build -t email-service .
	# rm email-service
run:
	# 设置了 GO_MICRO_BROKER 环境变量，go-micro 会使用该地址指定的 NATS 消息系统做事件的订阅和发布
	docker run --net="host" \
		-p 50054 \
		-e MICRO_REGISTRY=mdns \
		-e MICRO_BROKER=nats \
		-e MICRO_BROKER_ADDRESS=0.0.0.0:4222 \
		email-service