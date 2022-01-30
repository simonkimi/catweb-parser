package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/rpc"
	"unsafe"
)

/*
struct ParseResult {
	char* data;
	int len;
};
*/
import "C"

type ParseResult C.struct_ParseResult

func main() {
}

//export ParseData
func ParseData(input *C.char, size int) ParseResult {
	dest := make([]byte, size)
	copy(dest, (*(*[1 << 30]byte)(unsafe.Pointer(input)))[:size:size])
	marshal := parseData(dest)
	return ParseResult{
		data: (*C.char)(C.CBytes(marshal)),
		len:  C.int(len(marshal)),
	}
}

func parseData(input []byte) []byte {
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
