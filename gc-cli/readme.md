# grpc

* grpc protobuf cli 를 사용하여 자동으로 코드를 생성해주는 도와주는 프로그램

# official documents
1) https://protobuf.dev/reference/go/go-generated/
2) https://grpc.io/docs/languages/go/generated-code/
3) https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 사용법

1. protoc 설치
   https://github.com/protocolbuffers/protobuf/releases
2. 들어가서 바이너리 다운받아서 $HOME/sdk/bin 에 추가
	- 경로는 예제이며 필요에 따라 변경하여 사용한다
3. protobuf 인스톨 하기
	* $GOPATH/bin 에 protoc extension 이 설치된다

~~~bash
# $GOPATH/bin/protoc-gen-go 프로그램 컴파일
go get google.golang.org/protobuf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# $GOPATH/bin/protoc-gen-go-grpc 프로그램 컴파일
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

~~~

4. cli 사용하기

~~~bash
# $GOPATH/bin/go-grpc 프로그램 컴파일
go install github.com/d3v-friends/go-grpc/gc-cli@latest

# 사용
gc-cli --config grpc.yml
~~~

## 프로파일링

~~~bash
go get github.com/google/pprof
~~~
