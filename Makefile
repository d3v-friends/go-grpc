run:
	go run github.com/d3v-friends/go-grpc/cli protoc --config=./sample/protoc.yaml
build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -ldflags '-s' -o bin/go-gprc.arm64 cli/main.go ;
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-s' -o bin/go-grpc.amd64 cli/main.go;
