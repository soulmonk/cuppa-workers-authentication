FROM golang

RUN apt-get update && apt-get -y upgrade && \
 apt-get -y install protobuf-compiler

# Create app directory
WORKDIR /usr/src/app

# easy rebuild?
COPY third_party/protoc-gen.sh ./third_party/protoc-gen.sh
COPY go.mod Makefile ./

RUN make install

# Bundle app source
COPY . .

RUN make build-all

CMD [ "make", "run-server" ]