package selector

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"golang.org/x/net/html"
	"strings"
)

func FindElement(p *protobuf.Selector, root *html.Node) []*html.Node {
	var node []*html.Node

	if strings.HasPrefix(p.Selector, "/") { // XPath选择器
		node = htmlquery.Find(root, p.Selector)
	} else {
		document := goquery.NewDocumentFromNode(root)
		node = document.Find(p.Selector).Nodes
	}

	return node
}

func callQuery(p *protobuf.Selector, root *html.Node) (string, error) {
	selector := p.Selector

	if selector == "" {
		if !(p.Function != protobuf.SelectorFunction_SELECTOR_FUNCTION_AUTO || p.Param != "" || p.Regex != "" || p.Replace != "" || p.Js != "") {
			if p.DefaultValue != "" {
				return p.DefaultValue, nil
			}
			return "", nil
		}
	}

	node := FindElement(p, root)[0]

	if node == nil {
		return "", errors.New("no nodes")
	}

	switch p.Function {
	case protobuf.SelectorFunction_SELECTOR_FUNCTION_ATTR:
		for _, key := range strings.Split(p.Param, ",") {
			start := strings.TrimSpace(key)
			return htmlquery.SelectAttr(node, start), nil
		}
	case protobuf.SelectorFunction_SELECTOR_FUNCTION_RAW, protobuf.SelectorFunction_SELECTOR_FUNCTION_AUTO:
		return htmlquery.OutputHTML(node, true), nil
	case protobuf.SelectorFunction_SELECTOR_FUNCTION_TEXT:
		return htmlquery.InnerText(node), nil
	}
	return "", nil
}

func ParseElement(p *protobuf.Selector, root *html.Node) (string, error) {
	function, err := callQuery(p, root)
	if err != nil {
		return "", nil
	}

	function, err = CallReg(p, function)
	if err != nil {
		return "", nil
	}

	function, err = CallJs(p, function)
	if err != nil {
		return "", nil
	}

	return function, nil
}
