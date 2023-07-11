//go:generate protoc --go_out=. --go-grpc_out=. ./user.proto
//go:generate protoc --go_out=. --go-grpc_out=. ./subscription.proto
//go:generate protoc --go_out=. --go-grpc_out=. ./settings.proto
package inboxapi
