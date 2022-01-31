package main

import (
	"github.com/simonkimi/catweb-parser/rpc"
	"unsafe"
)

/*
#include <stdlib.h>
struct ParseResult {
	char* data;
	int len;
};
*/
import "C"

type ParseResult C.struct_ParseResult

//export ParseData
func ParseData(input *C.char, size C.int) ParseResult {
	dest := C.GoBytes(unsafe.Pointer(input), size)
	marshal := rpc.ExecSelector(dest)
	return ParseResult{
		data: (*C.char)(C.CBytes(marshal)),
		len:  C.int(len(marshal)),
	}
}

//export FreeResult
func FreeResult(result ParseResult) {
	C.free(unsafe.Pointer(result.data))
}

func main() {
}
