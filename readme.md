## 사용방법

1. protoc 설치
    * linux 계열 (darwin 포함)

~~~bash
   # 버전은 github 방문하여 주소 복사해오기
   # https://github.com/protocolbuffers/protobuf/releases/download/v25.0/protoc-25.0-linux-aarch_64.
   PROTOC_VERS=25_0
   PROTOC_FILE="protoc_$PROTOC_VERS.zip"
   PROTOC_PATH="$HOME/sdk/protoc/$PROTOC_VERS"
   wget https://github.com/protocolbuffers/protobuf/releases/download/v25.0/protoc-25.0-linux-aarch_64.zip -O "$PROTOC_FILE"
   mkdir -p "$PROTOC_PATH"
   unzip $PROTOC_FILE -d "$PROTOC_PATH"

   echo "set env to .bashrc"
   echo "# PROTOC" >> .bashrc
   echo "PROTOC_VERS=$PROTOC_VERS" >> .bashrc 
   echo "PATH=\$PATH:\$HOME/sdk/protoc/$PROTOC_VERS/bin" >> .bashrc

   rm "$PROTOC_FILE"
   /bin/bash
~~~

2. windows
    * https://github.com/protocolbuffers/protobuf 들어가서 바이너리 다운받아서 $HOME/bin 에 추가

---

3. protobuf 인스톨 하기
    * $GOPATH/bin 에 protoc extension 이 설치된다

~~~bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
~~~

---

4. tools.go 추가하기

* 프로젝트 root path 에서 실행힌다

~~~bash
echo "//go:build tools
// +build tools

package tools

import (
	_ \"github.com/d3v-friends/go-grpc\"
)
" >> tools.go
~~~

5. 사용

~~~bash
go run github.com/d3v-friends/go-grpc protoc --config=config.yaml
~~~
