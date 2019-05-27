build:
	GOOS=linux GOARCH=amd64 go build -o user-cli
	docker build -t user-cli .
	rm user-cli
run:
	docker run -e MICRO_REGISTRY=mdns user-cli