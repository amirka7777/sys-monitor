server:
	go run cmd/server/main.go
agent:
	go run cmd/agent/main.go
proto:
	protoc --go_out=. --go--options=paths=source_relative --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/metrics.proto