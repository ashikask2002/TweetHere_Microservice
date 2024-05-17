run:
	go run cmd/main.go	

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth.proto

prototweet:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/tweet/tweet.proto 