# gRPC for DEV

## 安裝

```bash
# 安裝 compiler，安裝完後就會有 protoc CLI 工具
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+

# 安裝 grpc-go
$ go get -u google.golang.org/grpc

# 安裝 protoc-gen-go
$ go get github.com/golang/protobuf/protoc-gen-go    # v1.22.0
$ protoc-gen-go --version      # 檢視版本
```

## Compile Proto 檔

```bash
# 進到有 proto 檔的資料夾
$ cd internal/grpc/proto/jubox

# compile 出 .pb.go 檔
$ protoc --go_out=plugins=grpc:. *.proto --go_opt=paths=source_relative
```