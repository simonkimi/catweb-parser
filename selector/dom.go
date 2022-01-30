package selector

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"golang.org/x/net/html"
)

type DomSelector struct {
	Env map[string]string
}

func (p *DomSelector) Nodes(selector *protobuf.Selector, parent *html.Node) []*html.Node {
	return FindElement(selector, parent)
}

func (p *DomSelector) String(selector *protobuf.Selector, parent *html.Node) string {
	result, err := ParseElement(selector, parent)
	if err != nil {
		return "err" + err.Error()
	}
	return utils.EnvReplace(result, p.Env)
}

func (p *DomSelector) Int(selector *protobuf.Selector, parent *html.Node) int32 {
	result, err := ParseElement(selector, parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.Env)
	return int32(utils.String2Int(input))
}

func (p *DomSelector) Double(selector *protobuf.Selector, parent *html.Node) float64 {
	result, err := ParseElement(selector, parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.Env)
	return utils.String2Double(input)
}

func (p *DomSelector) Color(selector *protobuf.Selector, parent *html.Node) *protobuf.ColorRpcModel {
	result, err := ParseElement(selector, parent)
	if err != nil {
		return nil
	}
	input := utils.EnvReplace(result, p.Env)
	return utils.String2Color(input)
}

func (p *DomSelector) GetEnv(extras []*protobuf.ExtraSelector, filter protobuf.ExtraSelectorType, parent *html.Node) (map[string]string, map[string]string) {
	local := make(map[string]string)
	global := make(map[string]string)

	for _, extra := range extras {
		if extra.Type == filter {
			if extra.Global {
				global[extra.Id] = p.String(extra.Selector, parent)
			} else {
				local[extra.Id] = p.String(extra.Selector, parent)
			}
		}
	}
	return local, global
}
