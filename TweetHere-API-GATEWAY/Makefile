run:
	go run cmd/main.go	

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/auth/auth.proto

prototweet:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/tweet/tweet.proto 

protochat:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/chat/chat.proto

protonoti:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/noti/noti.proto

swag:
	swag init -g cmd/main.go -o ./cmd/docs

