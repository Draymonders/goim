package api

//go:generate protoc -I. -I$GOPATH/src/github.com/Terry-Mao/goim/api --go_out=plugins=grpc:. --go_opt=paths=source_relative protocol/protocol.proto
//go:generate protoc -I. -I$GOPATH/src/github.com/Terry-Mao/goim/api --go_out=plugins=grpc:. --go_opt=paths=source_relative comet/comet.proto
//go:generate protoc -I. -I$GOPATH/src/github.com/Terry-Mao/goim/api --go_out=plugins=grpc:. --go_opt=paths=source_relative logic/logic.proto
