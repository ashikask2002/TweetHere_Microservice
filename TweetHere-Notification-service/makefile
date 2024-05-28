run:
	go run cmd/main.go	

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/noti/noti.proto

protoauth:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth.proto