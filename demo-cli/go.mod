module github.com/cg-0508/laracom/demo-cli

go 1.13

require (
	github.com/cg-0508/laracom/demo-service v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/protoc-gen-micro v1.0.0
)

replace github.com/cg-0508/laracom/demo-service => /Users/chengang/code/laracom/demo-service

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
