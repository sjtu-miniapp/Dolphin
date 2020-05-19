module github.com/sjtu-miniapp/dolphin

go 1.13

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.4 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/prometheus/client_golang v1.3.0 // indirect
	github.com/rs/cors v1.7.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20200501053045-e0ff5e5a1de5 // indirect
	golang.org/x/tools v0.0.0-20200103221440-774c71fcf114 // indirect
	google.golang.org/genproto v0.0.0-20200430143042-b979b6f78d84 // indirect
	gopkg.in/yaml.v2 v2.2.8
)
