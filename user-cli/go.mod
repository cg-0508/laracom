module github.com/cg-0508/laracom/user-cli

go 1.13

require (
	github.com/cg-0508/laracom/user-service v0.0.0-20201128081002-4fa4064c1238
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

//replace github.com/cg-0508/laracom/user-service => /Users/chengang/code/laracom/user-service
