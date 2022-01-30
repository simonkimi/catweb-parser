package main

import "C"
import (
	"github.com/golang/protobuf/proto"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/rpc"
)

func main() {

}

//export ParseData
func ParseData(input []byte) []byte {
	var errmsg string
	data, err := rpc.ParseRpcData(input)
	if err != nil {
		errmsg = err.Error()
	} else {
		errmsg = ""
	}
	model := &protobuf.RpcResponse{
		Data:  data,
		Error: errmsg,
	}
	marshal, err := proto.Marshal(model)
	if err != nil {
		panic(err)
	}
	return marshal
}
