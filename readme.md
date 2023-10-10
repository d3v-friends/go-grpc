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
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ~~~

3. cli 설치w
   ~~~bash
   go install github.com/d3v-friends/go-grpc
   ~~~

4. 사용
   ~~~bash
   go-grpc protoc --config=protoc.yaml
   ~~~
