echo "Generating gRPC stubs..."
protoc --go_out=. --go_opt=paths=import --go_opt=Mgo-memps-proto/publisher/v1/service.proto=pkg/publisher/proto --go-grpc_out=. --go-grpc_opt=paths=import --go-grpc_opt=Mgo-memps-proto/publisher/v1/service.proto=pkg/publisher/proto go-memps-proto/publisher/v1/*.proto
protoc --go_out=. --go_opt=paths=import --go_opt=Mgo-memps-proto/subscriber/v1/service.proto=pkg/subscriber/proto --go-grpc_out=. --go-grpc_opt=paths=import --go-grpc_opt=Mgo-memps-proto/subscriber/v1/service.proto=pkg/subscriber/proto go-memps-proto/subscriber/v1/*.proto
echo "Done"
sleep 3
