package utils

import (
	"github.com/robertkrimen/otto"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"strconv"
	"strings"
)

func String2Int(input string) int64 {
	parseInt, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		return parseInt
	}

	vm := otto.New()
	runString, err := vm.Run(input)
	if err != nil {
		return 0
	}
	val, err := runString.ToInteger()
	if err != nil {
		return 0
	}
	return val
}

func String2Double(input string) float64 {
	result, err := strconv.ParseFloat(input, 64)
	if err == nil {
		return result
	}
	vm := otto.New()
	runString, err := vm.Run(input)
	if err != nil {
		return 0.0
	}
	val, err := runString.ToFloat()
	if err != nil {
		return 0.0
	}
	return val
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

func Map[T any, E any](input []T, function func(item T) E) []E {
	var list []E
	for _, item := range input {
		list = append(list, function(item))
	}
	return list
}
