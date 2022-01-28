package main

import (
	"fmt"
	js "github.com/dop251/goja"
)

func main() {
	vm := js.New() // 创建engine实例
	result, _ := vm.RunString("1+1")
	data := result.Export()
	fmt.Printf("Result %d", data)
}
