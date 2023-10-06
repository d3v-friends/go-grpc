run:
	go run ./cli/main.go protoc --config ./sample/config.yaml;

runApp:
	go-grpc-cli protoc  --config ./sample/config.yaml;
build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o ./bin/go-gprc.exe cli.go;
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -ldflags '-s' -o ./bin/go-gprc.arm64 cli.go;
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-s' -o ./bin/go-grpc.amd64 cli.go;
