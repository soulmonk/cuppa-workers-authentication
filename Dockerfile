FROM golang

RUN apt-get update && apt-get -y upgrade
RUN apt-get -y install autoconf automake libtool curl make g++ unzip  protobuf-compiler

# Create app directory
WORKDIR /usr/src/app

# Bundle app source
COPY . .

#RUN go mod cuppa-workers-authentication
RUN make init
RUN make go-mod-init
RUN make build-all

CMD [ "make", "run-server" ]