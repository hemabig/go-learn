1. 安装 protoc; 下载 压缩包, 把protoc 复制到 bin 目录
2. 安装 protoc-gen-go 和 protoc-gen-go-grpc, 指定 bin 目录
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
3. protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./grpc.proto
