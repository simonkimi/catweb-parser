package parser

import (
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
