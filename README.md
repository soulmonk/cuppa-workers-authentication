# Authentication service
Install packages
- `make install`

Build application
- `make buil-all`

Run application
- `make run-server`

## Requirements

- Generate code for grpc `protobuf` is required, install it with `brew`:
  ```shell
  brew install protobuf protoc-gen-go protoc-gen-go-grpc
  ```

## TODO
- use go mod
- ??? user id in login response

## docker
- `docker run -it -v `pwd`:/app -w /app golang bash` in power shell`${PWD}` instead `pwd`
- `docker build -t cuppa-workers-authentication .`
- `docker build -t cuppa/workers-authentication:v1 --platform linux/arm64 .`
- `docker tag cuppa/workers-authentication:v1 rpisoulv1.kube:31320/cuppa/workers-authentication:latest`
- `docker push rpisoulv1.kube:31320/cuppa/workers-authentication:latest`

- `docker run -d --name cuppa-workers-authentication -p 9090:9090 -e PORT=${PORT} cuppa-${SERVICE_NAME}`
