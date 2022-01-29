package utils

import (
	"github.com/dop251/goja"
	"strconv"
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
