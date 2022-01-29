package parser

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"golang.org/x/net/html"
)

type DomSelector struct {
	env map[string]string
}

func (p *DomSelector) Nodes(selector *Selector, parent *html.Node) []*html.Node {
	return selector.FindElement(parent)
}

func (p *DomSelector) String(selector *Selector, parent *html.Node) string {
	result, err := selector.Find(parent)
	if err != nil {
		return "err" + err.Error()
	}
	return utils.EnvReplace(result, p.env)
}

func (p *DomSelector) Int(selector *Selector, parent *html.Node) int {
	result, err := selector.Find(parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.env)
	return int(utils.String2Int(input))
}

func (p *DomSelector) Double(selector *Selector, parent *html.Node) float64 {
	result, err := selector.Find(parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.env)
	return utils.String2Double(input)
}

func (p *DomSelector) Color(selector *Selector, parent *html.Node) *protobuf.ColorPb {
	result, err := selector.Find(parent)
	if err != nil {
		return nil
	}
	input := utils.EnvReplace(result, p.env)
	return utils.String2Color(input)
}
