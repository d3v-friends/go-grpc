## 사용방법

1. protoc 설치
   ~~~bash
   brew install protobuf
   ~~~
   
2. protobuf 인스톨 하기
   ~~~bash
   go get go get google.golang.org/protobuf/cmd/protoc-gen-go
   go install google.golang.org/protobuf/cmd/protoc-gen-go
   ~~~

3. cli 설치
   ~~~bash
   go install github.com/d3v-friends/go-grpc/go-grpc
   ~~~

4. 사용
   ~~~bash
   go-grpc protoc --config=protoc.yaml
   ~~~
