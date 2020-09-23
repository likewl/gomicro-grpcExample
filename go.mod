module gomicro-grpcExample

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.1
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
