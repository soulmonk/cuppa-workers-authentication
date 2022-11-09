protoc \
  --proto_path=api/proto/admin --proto_path=third_party \
  --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:pkg/api/admin \
  --go-grpc_out=require_unimplemented_servers=false:pkg/api/admin \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  api/proto/admin/*.proto

protoc \
  --proto_path=api/proto/authentication --proto_path=third_party \
  --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:pkg/api/authentication \
  --go-grpc_out=require_unimplemented_servers=false:pkg/api/authentication \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  api/proto/authentication/*.proto


#protoc --proto_path=api/proto/authentication --proto_path=third_party --go_out=plugins=grpc:pkg/api/authentication \
#  --go_opt=paths=source_relative \
#  --go-grpc_opt=paths=source_relative \
# api/proto/authentication/*.proto

#protoc --proto_path=api/proto/authentication --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/authentication \
#  api/proto/authentication/*.proto
#protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/authentication api/proto/authentication/*.proto
#
#protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 admin-service.proto
#protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 admin-service.proto

#
