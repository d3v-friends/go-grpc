## 사용방법

1. protoc 설치
   ~~~bash
   # darwin
   brew install protobuf

   # windows
   # https://github.com/protocolbuffers/protobuf/releases 들어가서 바이너라 다운받아서 $HOME/bin 에 추가
   ~~~
   
2. protobuf 인스톨 하기
   ~~~bash
   go get google.golang.org/protobuf/cmd/protoc-gen-go
   go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
   
   go install google.golang.org/protobuf/cmd/protoc-gen-go
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
   ~~~

3. cli 설치
   ~~~bash
   go install github.com/d3v-friends/go-grpc/go-grpc-cli
   ~~~

4. 사용
   ~~~bash
   go-grpc-cli protoc --config=protoc.yaml
   ~~~
