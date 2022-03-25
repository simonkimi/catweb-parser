package selector

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"golang.org/x/net/html"
	"math"
	"strings"
)

func innerText(n *html.Node) string {
	var output func(*bytes.Buffer, *html.Node)
	output = func(buf *bytes.Buffer, n *html.Node) {
		switch n.Type {
		case html.TextNode:
			buf.WriteString(n.Data)
			return
		case html.CommentNode:
			return
		case html.ElementNode:
			if n.Data == "br" {
				buf.WriteString("\n")
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			output(buf, child)
		}
	}

	var buf bytes.Buffer
	output(&buf, n)
	return buf.String()
}

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

func callQuery(p *protobuf.Selector, root *html.Node) (string, bool) {
	if p.Selector == "" {
		if !(p.Function != protobuf.SelectorFunction_SELECTOR_FUNCTION_AUTO || p.Param != "" || p.Regex != "" || p.Replace != "" || p.Script != "") {
			return "", false
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
					return result, true
				}
			}
		case protobuf.SelectorFunction_SELECTOR_FUNCTION_RAW, protobuf.SelectorFunction_SELECTOR_FUNCTION_AUTO:
			return htmlquery.OutputHTML(node, true), true
		case protobuf.SelectorFunction_SELECTOR_FUNCTION_TEXT:
			return innerText(node), true
		}
	}

	return "", false
}

func ParseElement(p *protobuf.Selector, root *html.Node) (string, error) {
	if p == nil {
		return "", nil
	}
	function, isFound := callQuery(p, root)
	if p.DefaultValue != "" && !isFound {
		return p.DefaultValue, nil
	}

	function, err := CallReg(p, function)
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
	function, isFound := callQuery(p, root)

	if p.DefaultValue != "" && !isFound {
		function = p.DefaultValue
	}

	function, err := CallReg(p, function)
	if err != nil {
		return "CallReg is nil", nil
	}

	function, err = CallJs(p, function)
	if err != nil {
		return "CallJs is nil", nil
	}

	return function, nil
}

func (p *DomSelector) htmlNodes(selector *protobuf.Selector, parent *html.Node) []*html.Node {
	return FindElement(selector, parent)
}

func (p *DomSelector) htmlString(selector *protobuf.Selector, parent *html.Node) string {
	result, err := ParseElement(selector, parent)
	if err != nil {
		return "err" + err.Error()
	}
	return utils.EnvReplace(result, p.Env)
}

func (p *DomSelector) htmlInt(selector *protobuf.Selector, parent *html.Node) int32 {
	result, err := ParseElement(selector, parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.Env)
	return int32(utils.String2Int(input))
}

func (p *DomSelector) htmlDouble(selector *protobuf.Selector, parent *html.Node) float64 {
	if selector == nil || selector.Selector == "" {
		if selector != nil && selector.DefaultValue != "" {
			return utils.String2Double(utils.EnvReplace(selector.DefaultValue, p.Env))
		} else {
			return math.NaN()
		}
	}
	result, err := ParseElement(selector, parent)
	if err != nil {
		return 0
	}
	input := utils.EnvReplace(result, p.Env)
	return utils.String2Double(input)
}

func (p *DomSelector) htmlColor(selector *protobuf.Selector, parent *html.Node) *protobuf.ColorRpcModel {
	result, err := ParseElement(selector, parent)
	if err != nil {
		return nil
	}
	input := utils.EnvReplace(result, p.Env)
	return utils.String2Color(input)
}

func (p *DomSelector) htmlTest(selector *protobuf.Selector, parent *html.Node) string {
	result, err := ParseTest(selector, parent)
	if err != nil {
		return "err" + err.Error()
	}
	input := utils.EnvReplace(result, p.Env)
	return input
}

func (p *DomSelector) htmlGetEnv(extras []*protobuf.ExtraSelector, filter protobuf.ExtraSelectorType, parent *html.Node) (map[string]string, map[string]string) {
	local := make(map[string]string)
	global := make(map[string]string)

	for _, extra := range extras {
		result := p.htmlString(extra.Selector, parent)
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

func (p *DomSelector) htmlLocalEnv(extras []*protobuf.ExtraSelector, filter protobuf.ExtraSelectorType, parent *html.Node) map[string]string {
	local := make(map[string]string)
	for _, extra := range extras {
		if extra.Type == filter {
			if !extra.Global {
				result := p.htmlString(extra.Selector, parent)
				if result != "" {
					local[extra.Id] = result
				}
			}
		}
	}
	return local
}
