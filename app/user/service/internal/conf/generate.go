// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2021/10/25

package conf

//go:generate protoc --proto_path=. --proto_path=./../../../../../third_party --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./conf.proto
