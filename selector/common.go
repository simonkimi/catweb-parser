package selector

import (
	"encoding/json"
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
		js := strings.TrimSpace(p.Js)
		if strings.HasPrefix(js, "{") && strings.HasSuffix(js, "}") {
			// Json格式数据, 进行替换
			data := make(map[string]string)
			err := json.Unmarshal([]byte(js), &data)
			if err == nil {
				val, exist := data[strings.TrimSpace(input)]
				if exist {
					return val, nil
				}
				return input, nil
			}
		}

		// 其他格式数据, 进行转换
		vm := otto.New()
		_, err := vm.Run(js)
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
