build:
	# protoc -I. --go_out=.:$(GOPATH)/src/shippy/user-service --micro_out=.:$(GOPATH)/src/shippy/user-service proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o user-service
	docker build -t user-service .
	# rm user-service
run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service