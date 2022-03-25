package selector

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"golang.org/x/net/html"
)

type DomSelector struct {
	Env map[string]string
}

func (p *DomSelector) Nodes(selector *protobuf.Selector, parent any) []any {
	switch parent.(type) {
	case *html.Node:
		var result []any
		for _, node := range p.htmlNodes(selector, parent.(*html.Node)) {
			result = append(result, node)
		}
		return result
	}
	return FindJsonElement(selector, parent)
}

func (p *DomSelector) String(selector *protobuf.Selector, parent any) string {
	switch parent.(type) {
	case *html.Node:
		return p.htmlString(selector, parent.(*html.Node))
	}
	return p.jsonString(selector, parent)
}

func (p *DomSelector) Int(selector *protobuf.Selector, parent any) int32 {
	switch parent.(type) {
	case *html.Node:
		return p.htmlInt(selector, parent.(*html.Node))
	}
	return p.jsonInt(selector, parent)
}

func (p *DomSelector) Double(selector *protobuf.Selector, parent any) float64 {
	switch parent.(type) {
	case *html.Node:
		return p.htmlDouble(selector, parent.(*html.Node))
	}
	return p.jsonDouble(selector, parent)
}

func (p *DomSelector) Color(selector *protobuf.Selector, parent any) *protobuf.ColorRpcModel {
	switch parent.(type) {
	case *html.Node:
		return p.htmlColor(selector, parent.(*html.Node))
	}
	return p.jsonColor(selector, parent)
}

func (p *DomSelector) GetEnv(extras []*protobuf.ExtraSelector, filter protobuf.ExtraSelectorType, parent any) (map[string]string, map[string]string) {
	switch parent.(type) {
	case *html.Node:
		return p.htmlGetEnv(extras, filter, parent.(*html.Node))
	}
	return p.jsonGetEnv(extras, filter, parent)
}

func (p *DomSelector) LocalEnv(extras []*protobuf.ExtraSelector, filter protobuf.ExtraSelectorType, parent any) map[string]string {
	switch parent.(type) {
	case *html.Node:
		return p.htmlLocalEnv(extras, filter, parent.(*html.Node))
	}
	return p.jsonLocalEnv(extras, filter, parent)
}
