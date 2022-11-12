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

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

# GO_PATH := ~/go
# PATH := $PATH:/$GO_PATH/bin

## Install:
install: init go-mod-init ## run "init" "go-mod-init"

go-mod-init: ## Download dependencies
	@echo " > Download dependencies"
	@go mod download
	@echo " NOTE: without this installation: event if in go.mod, could not gen proto"
	@go get google.golang.org/protobuf \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@export PATH="$PATH:$(go env GOPATH)/bin"

init: ## Simple initialization. Make `third_party/protoc-gen.sh` executable
	@echo " > Simple initialization"
	@echo " >> Make 'third_party/protoc-gen.sh' executable"
	@chmod +x $(PROJ_BASE)/third_party/protoc-gen.sh
	@-mkdir -p pkg/api/v1
	@-mkdir -p api/swagger/v1

#clean:
#	@echo "  >  Clean build files. Runs `go clean` internally."
#	@-$(MAKE) clean-build clean-proto go-clean
clean: ## Clean build files.
	@echo "  >  Clean build files."
	@-$(MAKE) clean-build clean-proto

clean-build:
	@echo "  >  Clean build"
	@-rm $(PROJ_BUILD_PATH)/* 2> /dev/null

clean-proto:
	@echo "  >  Clean proto"
	@-rm -f $(PROJ_BASE)/pkg/api/**/*.pb.* 2> /dev/null
	@echo

gen-proto: ## Generate proto
	@echo "  >  Generate proto"
	@echo "${YELLOW}Hint${RESET}: in some cases require export path"
	@echo "export GO_PATH := ~/go"
	@echo 'export PATH := $$PATH:$$GO_PATH/bin'
	@$(shell $(PROJ_BASE)/third_party/protoc-gen.sh)


gen-db: ## Generate go from sql
	@sqlc generate


build-all: ## Runs `gen-proto` `build-server` `c`
	@echo "  >  Build all"
	@-$(MAKE) gen-proto build-server

run-server: ## RUN_OPTIONS='-grpc-port= -http-port= -db-host= --db-port= -db-user= -db-password= -db-name= -log-level=-1 -log-time-format=2006-01-02T15:04:05.999999999Z07:00'
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

build-server: ## Build grpc-server
	@echo "  >  Build server"
	go build -o $(PROJ_BUILD_PATH)/server $(PROJ_BASE)/cmd/server/server.go

build-migration: ## Build migration
	@echo "  >  Build server"
	go build -o $(PROJ_BUILD_PATH)/migration $(PROJ_BASE)/cmd/migration.go

build-grpc-client: ## Build grpc client
	@echo " > Build Grpc Client"
	 go build -o $(PROJ_BUILD_PATH)/client-grpc $(PROJ_BASE)/cmd/client-grpc/main.go

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get $(get)

show-up-deps: ## view available dependency upgrades
	@go list -u -m all

go-update-mod-version: ### Update go version
	@echo " > Update go version to $(go_version)"
	@go mod edit -go $(go_version)

## Docker:
install-swagger-docker: ## install swagger-docker
	@echo "  > Installing swagger docker"
	@docker pull swaggerapi/swagger-ui

## [[depricated]]
create-swagger-docker: ## create docker container "authentication-swagger-ui"
	@docker create --name authentication-swagger-ui -p 9000:8080 -e SWAGGER_JSON=/app/api/swagger/v1/authenctication-service.swagger.json -v ${PROJ_BASE}:/app swaggerapi/swagger-ui

## [[depricated]]
run-swagger-docker: ## running swagger docker
	@echo "  >  Starting swagger docker http://127.0.0.1:9000"
	@docker run --rm -p 9000:8080 -e SWAGGER_JSON=/app/api/swagger/authentication/openapi.json -v ${PROJ_BASE}:/app swaggerapi/swagger-ui

start-swagger-docker: ## starting swagger docker
	@echo "  >  Starting swagger docker http://127.0.0.1:9000"
	@docker start authentication-swagger-ui

stop-swagger-docker: ## stopping swagger docker
	@echo "  >  Stopping swagger docker"
	@docker stop authentication-swagger-ui

#go-clean:
#	@echo "  >  Cleaning build cache"
#	@go clean $(PROJ_BUILD_PATH)

.DEFAULT_GOAL := help
.PHONY: help
all: help
## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-24s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
