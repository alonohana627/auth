#!/bin/sh

protoc --proto_path=/AuthService/pb --go_out=/AuthService/pb /AuthService/pb/user.proto --experimental_allow_proto3_optional
protoc --proto_path=/AuthService/pb --go_out=/AuthService/pb /AuthService/pb/routes.proto --experimental_allow_proto3_optional
protoc --proto_path=/AuthService/pb --go-grpc_out=/AuthService/pb /AuthService/pb/routes.proto --experimental_allow_proto3_optional