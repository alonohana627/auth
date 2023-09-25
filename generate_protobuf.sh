#!/bin/sh

protoc --proto_path=/AuthService/pb --go_out=/AuthService/pb /AuthService/pb/user.proto
protoc --proto_path=/AuthService/pb --go_out=/AuthService/pb /AuthService/pb/routes.proto
protoc --proto_path=/AuthService/pb --go-grpc_out=/AuthService/pb /AuthService/pb/routes.proto