package test

import (
	"fmt"
	"github.com/ohler55/ojg/oj"
	"io/ioutil"
	"testing"
)

func TestJson(t *testing.T) {
	buffer, _ := ioutil.ReadFile("./suggest.json")
	root, err := oj.Parse(buffer)
	if err != nil {
		panic(err)
	}

	switch root.(type) {
	case map[string]interface{}:
		fmt.Println(oj.JSON(root))
	}
}
