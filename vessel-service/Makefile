build:
	protoc -I. --go_out=.:$(GOPATH)/src/shippy/vessel-service --micro_out=.:$(GOPATH)/src/shippy/vessel-service proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build -o vessel-service
	docker build -t vessel-service .
	rm vessel-service
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service