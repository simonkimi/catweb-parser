package selector

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"golang.org/x/net/html"
	"strings"
)

func FindElement(p *protobuf.Selector, root *html.Node) []*html.Node {
	if root == nil || p == nil {
		return []*html.Node{}
	}

	if p.Selector == "" {
		return []*html.Node{root}
	}

	if strings.HasPrefix(p.Selector, "/") {
		// XPath选择器
		result, err := htmlquery.QueryAll(root, p.Selector)
		if err != nil {
			return []*html.Node{}
		}
		return result
	} else {
		// Html选择器
		document := goquery.NewDocumentFromNode(root)
		return document.Find(p.Selector).Nodes
	}
}

func callQuery(p *protobuf.Selector, root *html.Node) (string, error, bool) {
	if p.Selector == "" {
		if !(p.Function != protobuf.SelectorFunction_SELECTOR_FUNCTION_AUTO || p.Param != "" || p.Regex != "" || p.Replace != "" || p.Js != "") {
			return "", nil, false
		}
	}

	nodes := FindElement(p, root)

	for _, node := range nodes {
		switch p.Function {
		case protobuf.SelectorFunction_SELECTOR_FUNCTION_ATTR:
			for _, key := range strings.Split(p.Param, ",") {
				start := strings.TrimSpace(key)
				result := htmlquery.SelectAttr(node, start)
				if result != "" {
					return result, nil, true
				}
			}
		case protobuf.SelectorFunction_SELECTOR_FUNCTION_RAW, protobuf.SelectorFunction_SELECTOR_FUNCTION_AUTO:
			return htmlquery.OutputHTML(node, true), nil, true
		case protobuf.SelectorFunction_SELECTOR_FUNCTION_TEXT:
			return htmlquery.InnerText(node), nil, true
		}
	}

	return "", nil, false
}

func ParseElement(p *protobuf.Selector, root *html.Node) (string, error) {
	if p == nil {
		return "", nil
	}
	function, err, isFound := callQuery(p, root)
	if p.DefaultValue != "" && !isFound {
		return p.DefaultValue, nil
	}
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

func ParseTest(p *protobuf.Selector, root *html.Node) (string, error) {
	if p == nil {
		return "Selector is nil", nil
	}
	function, err, isFound := callQuery(p, root)
	if err != nil {
		return "CallQuery is nil", nil
	}

	if p.DefaultValue != "" && !isFound {
		function = p.DefaultValue
	}

	function, err = CallReg(p, function)
	if err != nil {
		return "CallReg is nil", nil
	}

	function, err = CallJs(p, function)
	if err != nil {
		return "CallJs is nil", nil
	}

	return function, nil
}
