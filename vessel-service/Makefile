build:
	protoc -I. --go_out=.:$(GOPATH)/src/shippy/vessel-service --micro_out=.:$(GOPATH)/src/shippy/vessel-service proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build -o vessel-service
	docker build -t vessel-service .
	rm vessel-service
run:
	docker run --net="host" \
		-e MICRO_REGISTRY=mdns \
		vessel-service