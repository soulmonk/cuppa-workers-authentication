FROM golang:1.16.7-alpine as build

RUN apk update && apk add bash make protoc && rm -rf /var/cache/apk/*

RUN PATH="$GOPATH/bin:$PATH"

# Create app directory
WORKDIR /usr/src/app

# easy rebuild?
COPY third_party/protoc-gen.sh ./third_party/protoc-gen.sh
COPY go.mod go.sum Makefile ./

RUN make clean
RUN make install

# Bundle app source
COPY . .

RUN make build-all

#FROM gcr.io/distroless/base-debian10
#WORKDIR /
#COPY --from=build /usr/src/app/build/server /server
#USER nonroot:nonroot
#CMD [ "/server" ]

FROM golang:1.16.7-alpine

WORKDIR /usr/src/app/

COPY --from=build /usr/src/app/build/server /usr/src/app/server

CMD [ "/usr/src/app/server" ]
