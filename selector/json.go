package selector

import (
	"github.com/ohler55/ojg/jp"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"math"
	"strconv"
)

func FindJsonElement(p *protobuf.Selector, root any) []any {
	if root == nil {
		return []any{}
	}
	if p.Selector == "" {
		return []any{root}
	}
	path, err := jp.ParseString(p.Selector)
	if err != nil {
		return []any{}
	}

	return path.Get(root)
}

func callJsQuery(p *protobuf.Selector, root any) (string, bool) {
	if p.Selector == "" {
		if !(p.Function != protobuf.SelectorFunction_SELECTOR_FUNCTION_AUTO || p.Param != "" || p.Regex != "" || p.Replace != "" || p.Script != "") {
			return "", false
		}
	}
	nodes := FindJsonElement(p, root)

	if len(nodes) == 0 {
		return "", false
	}

	switch nodes[0].(type) {
	case string:
		return nodes[0].(string), true
	case int:
		return strconv.Itoa(nodes[0].(int)), true
	case float64:
		return strconv.FormatFloat(nodes[0].(float64), 'f', -1, 64), true
	case bool:
		return strconv.FormatBool(nodes[0].(bool)), true
	default:
		return "", false
	}
}

func ParseJsonElement(p *protobuf.Selector, root any) (string, error) {
	if p == nil {
		return "", nil
	}
	result, isFound := callJsQuery(p, root)
	if p.DefaultValue != "" && !isFound {
		return p.DefaultValue, nil
	}
	result, err := CallReg(p, result)
	if err != nil {
		return "", nil
	}
	result, err = CallJs(p, result)
	if err != nil {
		return "", nil
	}
	return result, nil
}

func (p *DomSelector) jsonNodes(selector *protobuf.Selector, parent any) []any {
	return FindJsonElement(selector, parent)
}

func (p *DomSelector) jsonString(selector *protobuf.Selector, parent any) string {
	result, err := ParseJsonElement(selector, parent)
	if err != nil {
		return "err" + err.Error()
	}
	return utils.EnvReplace(result, p.Env)
}

func (p *DomSelector) jsonInt(selector *protobuf.Selector, parent any) int32 {
	result, err := ParseJsonElement(selector, parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.Env)
	return int32(utils.String2Int(input))
}

func (p *DomSelector) jsonDouble(selector *protobuf.Selector, parent any) float64 {
	if selector == nil || selector.Selector == "" {
		if selector != nil && selector.DefaultValue != "" {
			return utils.String2Double(utils.EnvReplace(selector.DefaultValue, p.Env))
		} else {
			return math.NaN()
		}
	}
	result, err := ParseJsonElement(selector, parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.Env)
	return utils.String2Double(input)
}

func (p *DomSelector) jsonColor(selector *protobuf.Selector, parent any) *protobuf.ColorRpcModel {
	result, err := ParseJsonElement(selector, parent)
	if err != nil {
		return nil
	}
	input := utils.EnvReplace(result, p.Env)
	return utils.String2Color(input)
}

func (p *DomSelector) jsonTest(selector *protobuf.Selector, parent any) string {
	result, err := ParseJsonElement(selector, parent)
	if err != nil {
		return "err" + err.Error()
	}
	input := utils.EnvReplace(result, p.Env)
	return input
}

func (p *DomSelector) jsonGetEnv(extras []*protobuf.ExtraSelector, filter protobuf.ExtraSelectorType, parent any) (map[string]string, map[string]string) {
	local := make(map[string]string)
	global := make(map[string]string)

	for _, extra := range extras {
		result := p.jsonString(extra.Selector, parent)
		if result != "" {
			if extra.Type == filter {
				if extra.Global {
					global[extra.Id] = result
				} else {
					local[extra.Id] = result
				}
			}
		}
	}
	return local, global
}

func (p *DomSelector) jsonLocalEnv(extras []*protobuf.ExtraSelector, filter protobuf.ExtraSelectorType, parent any) map[string]string {
	local := make(map[string]string)
	for _, extra := range extras {
		if extra.Type == filter {
			if !extra.Global {
				result := p.jsonString(extra.Selector, parent)
				if result != "" {
					local[extra.Id] = result
				}
			}
		}
	}
	return local
}
