build:
	protoc --proto_path=. --micro_out=. --go_out=. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64
	docker build -t laracom-demo-service .
run:
	docker run -d -p 9091:9091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns laracom-demo-service