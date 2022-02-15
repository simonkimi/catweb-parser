package selector

import (
	"github.com/robertkrimen/otto"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"regexp"
	"strconv"
	"strings"
)

func CallReg(p *protobuf.Selector, input string) (string, error) {
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
			return match[len(match)-1], nil
		} else {
			rep := p.Replace
			for i := len(match) - 1; i >= 1; i-- {
				rep = strings.Replace(rep, "$"+strconv.Itoa(i), match[i], -1)
			}
			return rep, nil
		}
	} else {
		return input, nil
	}
}

func CallJs(p *protobuf.Selector, input string) (string, error) {
	if p.Js != "" {
		vm := otto.New()
		_, err := vm.Run(p.Js)
		if err != nil {
			return "", err
		}

		runString, err := vm.Run("hook('" + input + "')")
		if err != nil {
			return "", err
		}

		return runString.String(), nil

	} else {
		return input, nil
	}
}
