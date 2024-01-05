run:
	go run ./cli.go protoc --config ./sample/config.yaml;
runApp:
	go run github.com/d3v-friends/go-grpc protoc --config ./sample/config.yaml;
build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o ./bin/go-grpc.exe cli.go;
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -ldflags '-s' -o ./bin/go-grpc.arm64 cli.go;
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-s' -o ./bin/go-grpc.amd64 cli.go;
protoc:
	go run ./cli.go protoc --config=config.yaml;
	git add ./protoc --all;


