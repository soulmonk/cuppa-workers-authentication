#-include .env

PROJECT_NAME := $(shell basename "$(PWD)")

PROJ_BASE := $(shell pwd -LP)
PROJ_BUILD_PATH := $(PROJ_BASE)/build

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECT_NAME)-stderr.txt
# Redirect error output to a file, so we can show it in development mode.
STDOUT := /tmp/.$(PROJECT_NAME)-stdout.txt

## install: init
install: init go-mod-init

## go-mod-init: Download dependencies
go-mod-init:
	@echo " > Download dependencies"
	@go mod download
	@echo " NOTE: without this installation: event if in go.mod, could not gen proto"
	@go get github.com/golang/protobuf/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@export PATH="$PATH:$(go env GOPATH)/bin"

## init: Simple initialization. Make `third_party/protoc-gen.sh` executable
init:
	@echo " > Simple initialization"
	@echo " >> Make 'third_party/protoc-gen.sh' executable"
	@chmod +x $(PROJ_BASE)/third_party/protoc-gen.sh
	@-mkdir -p pkg/api/v1
	@-mkdir -p api/swagger/v1

## clean: Clean build files.
#clean:
#	@echo "  >  Clean build files. Runs `go clean` internally."
#	@-$(MAKE) clean-build clean-proto go-clean
clean:
	@echo "  >  Clean build files."
	@-$(MAKE) clean-build clean-proto

clean-build:
	@echo "  >  Clean build"
	@-rm $(PROJ_BUILD_PATH)/* 2> /dev/null

clean-proto:
	@echo "  >  Clean proto"
	@-rm -f $(PROJ_BASE)/api/swagger/**/*.json 2> /dev/null
	@-rm -f $(PROJ_BASE)/pkg/api/v1/*.pb.* 2> /dev/null
	@echo

## gen-proto: Generate proto
gen-proto:
	@echo "  >  Generate proto"
	@$(shell $(PROJ_BASE)/third_party/protoc-gen.sh)

## build-all: Runs `gen-proto` `build-server` `c`
build-all:
	@echo "  >  Build all"
	@-$(MAKE) gen-proto build-server

## run-server: RUN_OPTIONS='-grpc-port= -http-port= -db-host= --db-port= -db-user= -db-password= -db-name= -log-level=-1 -log-time-format=2006-01-02T15:04:05.999999999Z07:00'
run-server:
	@echo "  >  Running server"
	@$(PROJ_BUILD_PATH)/server $(RUN_OPTIONS)

start-server: stop-server
	@echo "  >  start server"
	@-$(PROJ_BUILD_PATH)/server $(RUN_OPTIONS) 2>$(STDOUT) & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"
	@echo "  >  stoud at $(STDOUT)"

stop-server:
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)


logs:
	@tail -f -n 100 $(STDOUT)

## build-server: Build grpc-server
build-server:
	@echo "  >  Build server"
	go build -o $(PROJ_BUILD_PATH)/server $(PROJ_BASE)/cmd/server/server.go

## build-migration: Build migration
build-migration:
	@echo "  >  Build server"
	go build -o $(PROJ_BUILD_PATH)/migration $(PROJ_BASE)/cmd/migration.go

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get $(get)

## install-swagger-docker: install swagger-docker
install-swagger-docker:
	@echo "  > Installing swagger docker"
	@docker pull swaggerapi/swagger-ui

## create-swagger-docker: create docker container "authentication-swagger-ui"
create-swagger-docker:
	@docker create --name authentication-swagger-ui -p 9000:8080 -e SWAGGER_JSON=/app/api/swagger/v1/authenctication-service.swagger.json -v ${PROJ_BASE}:/app swaggerapi/swagger-ui

## run-swagger-docker: running swagger docker
run-swagger-docker:
	@echo "  >  Starting swagger docker http://127.0.0.1:9000"
	@docker run -p 9000:8080 -e SWAGGER_JSON=/app/api/swagger/v1/authenctication-service.swagger.json -v ${PROJ_BASE}:/app swaggerapi/swagger-ui

## start-swagger-docker: starting swagger docker
start-swagger-docker:
	@echo "  >  Starting swagger docker http://127.0.0.1:9000"
	@docker start authentication-swagger-ui

## stop-swagger-docker: stopping swagger docker
stop-swagger-docker:
	@echo "  >  Stopping swagger docker"
	@docker stop authentication-swagger-ui

#go-clean:
#	@echo "  >  Cleaning build cache"
#	@go clean $(PROJ_BUILD_PATH)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
