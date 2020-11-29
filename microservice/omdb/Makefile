gen :
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/omdb.proto
init :
	go mod vendor
run :
	go run main.go
test :
	go test ./... -cover
docker-build :
	docker build -t omdb:latest .
docker-run :
	docker run -it -p 3000:3000 -p 4000:4000 omdb:latest



