FROM golang:1.20

RUN apt update && apt-get install protobuf-compiler -y
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


WORKDIR /AuthService

COPY go.mod go.sum /AuthService/
RUN go mod download

COPY . /AuthService

RUN chmod +x /AuthService/generate_protobuf.sh
RUN /AuthService/generate_protobuf.sh
RUN go build .