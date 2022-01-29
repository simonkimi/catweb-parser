package parser

import (
	"github.com/dop251/goja"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"regexp"
	"strconv"
	"strings"
)

type Selector protobuf.Selector

func (p *Selector) CallReg(input string) (string, error) {
	if p.Regex != "" {
		reg, err := regexp.Compile(p.Regex)
		if err != nil {
			return "", err
		}
		match := reg.FindStringSubmatch(input)
		if match == nil {
			return "", nil
		}
		if p.Replace == "" {
			return match[0], err
		} else {
			rep := p.Replace
			for i := 1; i < len(match); i++ {
				rep = strings.Replace(rep, "$"+strconv.Itoa(i), match[1], -1)
			}
			return rep, nil
		}
	} else {
		return input, nil
	}
}

func (p *Selector) CallJs(input string) (string, error) {
	if p.Js != "" {
		vm := goja.New()
		_, err := vm.RunString(p.Js)
		if err != nil {
			return "", err
		}

		runString, err := vm.RunString("hook(" + input + ")")
		if err != nil {
			return "", err
		}

		return runString.Export().(string), nil

	} else {
		return input, nil
	}
}
