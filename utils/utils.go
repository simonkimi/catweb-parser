package utils

import (
	"regexp"
	"strings"
)

func EnvReplace(input string, env map[string]string) string {
	for k, v := range env {
		input = strings.ReplaceAll(input, "{"+k+"}", v)
	}
	reg, err := regexp.Compile(`\${(?P<var>\w+):(?P<context>[^:]*):?(?P<default>.*)}`)
	if err != nil {
		return input
	}
	matches := reg.FindAllStringSubmatch(input, -1)

	if matches == nil {
		return input
	}
	for _, match := range matches {
		group := make(map[string]string)
		for i, name := range reg.SubexpNames() {
			if i != 0 && name != "" {
				group[name] = match[i]
			}
		}
		varName, _ := group["var"]
		context, _ := group["context"]
		defaultValue, defaultExist := group["default"]
		if _, exist := env[varName]; exist {
			input = strings.Replace(input, match[0], context, -1)
		} else {
			if defaultExist {
				input = strings.Replace(input, match[0], defaultValue, -1)
			} else {
				input = strings.Replace(input, match[0], "", -1)
			}
		}
	}
	return input
}
