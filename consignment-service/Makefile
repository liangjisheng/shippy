build:
	protoc -I. --go_out=.:$(GOPATH)/src/shippy/consignment-service --micro_out=.:$(GOPATH)/src/shippy/consignment-service proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build -o consignment-service
	docker build -t consignment-service .
	rm consignment-service
run:
	docker run --net="host" \
		-e MICRO_REGISTRY=mdns \
		consignment-service