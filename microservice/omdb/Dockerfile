#build stage
FROM golang:1.15
ENV GO111MODULE=on
RUN apt-get update -y && \
     apt-get upgrade -y && \
     apt-get dist-upgrade -y && \
     apt-get -y autoremove && \
     apt-get clean
RUN apt-get install -y git protobuf-compiler
WORKDIR /go/src/app
COPY . .
RUN PATH="$PATH:$(go env GOPATH)/bin"
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN make init
RUN make gen
RUN go install -v ./...
ENTRYPOINT /go/bin/omdb
LABEL Name=omdb Version=0.0.1
EXPOSE 3000
EXPOSE 4000