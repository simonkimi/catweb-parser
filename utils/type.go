package utils

import (
	"github.com/dop251/goja"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"strconv"
	"strings"
)

func String2Int(input string) int64 {
	parseInt, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		return parseInt
	}

	vm := goja.New()
	runString, err := vm.RunString(input)
	if err != nil {
		return 0
	}
	return runString.ToInteger()
}

func String2Double(input string) float64 {
	parseInt, err := strconv.ParseFloat(input, 64)
	if err == nil {
		return parseInt
	}

	vm := goja.New()
	runString, err := vm.RunString(input)
	if err != nil {
		return 0
	}
	return runString.ToFloat()
}

func String2Color(input string) *protobuf.ColorRpcModel {
	var colorHex string
	if strings.HasPrefix(input, "0x") || strings.HasPrefix(input, "0X") {
		colorHex = input[2:]
	} else if strings.HasPrefix(input, "#") {
		colorHex = input[1:]
	}
	parseInt, err := strconv.ParseInt(colorHex, 16, 64)
	if err != nil {
		return nil
	}
	return &protobuf.ColorRpcModel{
		A: int32(parseInt >> 24 & 0xFF),
		R: int32(parseInt >> 16 & 0xFF),
		G: int32(parseInt >> 8 & 0xFF),
		B: int32(parseInt & 0xFF),
	}
}
