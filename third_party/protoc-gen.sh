protoc \
  --proto_path=api/proto/v1 --proto_path=third_party \
  --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:pkg/api/v1 \
  --go-grpc_out=require_unimplemented_servers=false:pkg/api/v1 \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  authenctication-service.proto

#protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 authenctication-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 authenctication-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 authenctication-service.proto

#
