module github.com/cg-0508/laracom/demo-service

go 1.13

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
