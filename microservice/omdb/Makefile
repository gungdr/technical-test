gen :
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/omdb.proto
init :
	go mod vendor
run :
	go run main.go
test :
	go test ./...


