package main

import (
	"github.com/robertkrimen/otto"
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

//export RunJs
func RunJs(js *C.char, input *C.char) ParseResult {
	jsStr := C.GoString(js)
	inputStr := C.GoString(input)
	vm := otto.New()
	_, err := vm.Run(jsStr)
	if err != nil {
		return ParseResult{
			data: (*C.char)(C.CString("")),
			len:  -1,
		}
	}
	runString, err := vm.Run("hook('" + inputStr + "')")
	if err != nil {
		return ParseResult{
			data: (*C.char)(C.CString("")),
			len:  -1,
		}
	}
	result := runString.String()
	return ParseResult{
		data: (*C.char)(C.CString(result)),
		len:  C.int(len(result)),
	}
}

func main() {
}
